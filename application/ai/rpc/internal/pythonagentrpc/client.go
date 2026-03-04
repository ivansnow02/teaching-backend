// Package pythonagentrpc 封装了对 Python AI Agent 的 HTTP 调用
package pythonagentrpc

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Client 是 Python Agent HTTP 客户端
type Client struct {
	baseUrl    string
	apiKey     string
	httpClient *http.Client
}

func NewClient(baseUrl, apiKey string, httpClient *http.Client) *Client {
	return &Client{
		baseUrl:    strings.TrimRight(baseUrl, "/"),
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

// ===================== 判卷 =====================

type GradeReq struct {
	RecordId        int64  `json:"record_id"`
	QuestionId      int64  `json:"question_id"`
	QuestionContent string `json:"question_content"`
	StandardAnswer  string `json:"standard_answer"`
	KnowledgePoints string `json:"knowledge_points"`
	UserAnswer      string `json:"user_answer"`
	MaxScore        string `json:"max_score"`
}

type GradeRes struct {
	QuestionId int64  `json:"question_id"`
	Score      string `json:"score"`
	IsCorrect  int32  `json:"is_correct"` // 0:错误 1:正确 2:部分
	AiComment  string `json:"ai_comment"`
}

func (c *Client) GradeSubjectiveQuestion(ctx context.Context, req *GradeReq) (*GradeRes, error) {
	var res GradeRes
	if err := c.postJSON(ctx, "/api/grade/subjective", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ===================== 课件向量化 =====================

type EmbedReq struct {
	MaterialId int64  `json:"material_id"`
	CourseId   int64  `json:"course_id"`
	Title      string `json:"title"`
	Url        string `json:"url"`
	Type       int32  `json:"type"`
}

type EmbedRes struct {
	Accepted bool `json:"accepted"`
}

func (c *Client) EmbedMaterial(ctx context.Context, req *EmbedReq) (*EmbedRes, error) {
	var res EmbedRes
	if err := c.postJSON(ctx, "/api/embed/material", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type EmbedStatusRes struct {
	AiStatus int32  `json:"ai_status"` // 0:未处理 1:处理中 2:已完成 3:失败
	Message  string `json:"message"`
}

func (c *Client) GetEmbedStatus(ctx context.Context, materialId int64) (*EmbedStatusRes, error) {
	var res EmbedStatusRes
	if err := c.getJSON(ctx, fmt.Sprintf("/api/embed/status?material_id=%d", materialId), &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ===================== 智能助教 RAG =====================

type AskReq struct {
	CourseId int64    `json:"course_id"`
	UserId   int64    `json:"user_id"`
	Question string   `json:"question"`
	History  []string `json:"history"`
}

type AskRes struct {
	Answer  string   `json:"answer"`
	Sources []string `json:"sources"`
}

func (c *Client) AskQuestion(ctx context.Context, req *AskReq) (*AskRes, error) {
	var res AskRes
	if err := c.postJSON(ctx, "/api/rag/ask", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// AskQuestionStreamChunk 是 SSE 流的一帧
type AskQuestionStreamChunk struct {
	Delta    string   `json:"delta"`
	Sources  []string `json:"sources"`
	Finished bool     `json:"finished"`
}

// AskQuestionStream 以 SSE 格式从 Python Agent 读取流式回答，通过 channel 逐帧传递
func (c *Client) AskQuestionStream(ctx context.Context, req *AskReq) (<-chan AskQuestionStreamChunk, <-chan error) {
	ch := make(chan AskQuestionStreamChunk, 8)
	errCh := make(chan error, 1)
	go func() {
		defer close(ch)
		defer close(errCh)
		if err := c.readSseStream(ctx, "/api/rag/ask/stream", req, func(data string) error {
			var chunk AskQuestionStreamChunk
			if err := json.Unmarshal([]byte(data), &chunk); err != nil {
				return err
			}
			select {
			case <-ctx.Done():
			case ch <- chunk:
			}
			return nil
		}); err != nil {
			errCh <- err
		}
	}()
	return ch, errCh
}

// ===================== 智能生成 (异步) =====================

type GenerateQuestionsReq struct {
	CourseId        int64  `json:"course_id"`
	KnowledgePoints string `json:"knowledge_points"`
	Count           int32  `json:"count"`
	Type            int32  `json:"type"`
	Difficulty      int32  `json:"difficulty"`
}

type GenerateCourseOutlineReq struct {
	CourseId     int64  `json:"course_id"`
	Topic        string `json:"topic"`
	Requirements string `json:"requirements"`
}

type GenerateCoursewareReq struct {
	ChapterId    int64  `json:"chapter_id"`
	Requirements string `json:"requirements"`
}

type GenerateTaskRes struct {
	TaskId string `json:"task_id"`
}

func (c *Client) GenerateQuestions(ctx context.Context, req *GenerateQuestionsReq) (*GenerateTaskRes, error) {
	var res GenerateTaskRes
	if err := c.postJSON(ctx, "/api/generate/questions", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GenerateCourseOutline(ctx context.Context, req *GenerateCourseOutlineReq) (*GenerateTaskRes, error) {
	var res GenerateTaskRes
	if err := c.postJSON(ctx, "/api/generate/course-outline", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GenerateCourseware(ctx context.Context, req *GenerateCoursewareReq) (*GenerateTaskRes, error) {
	var res GenerateTaskRes
	if err := c.postJSON(ctx, "/api/generate/courseware", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GenerateCoursewareStreamChunk 是课件生成 SSE 的一帧
type GenerateCoursewareStreamChunk struct {
	Delta    string `json:"delta"`
	Finished bool   `json:"finished"`
}

// GenerateCoursewareStream 以 SSE 格式流式读取 AI 生成的课件 Markdown
func (c *Client) GenerateCoursewareStream(ctx context.Context, req *GenerateCoursewareReq) (<-chan GenerateCoursewareStreamChunk, <-chan error) {
	ch := make(chan GenerateCoursewareStreamChunk, 8)
	errCh := make(chan error, 1)
	go func() {
		defer close(ch)
		defer close(errCh)
		if err := c.readSseStream(ctx, "/api/generate/courseware/stream", req, func(data string) error {
			var chunk GenerateCoursewareStreamChunk
			if err := json.Unmarshal([]byte(data), &chunk); err != nil {
				return err
			}
			select {
			case <-ctx.Done():
			case ch <- chunk:
			}
			return nil
		}); err != nil {
			errCh <- err
		}
	}()
	return ch, errCh
}

// ===================== 任务状态查询 =====================

type AiTaskStatusRes struct {
	Status  int32  `json:"status"` // 0:处理中 1:已完成 2:失败
	Result  string `json:"result"`
	Message string `json:"message"`
}

func (c *Client) GetAiTaskStatus(ctx context.Context, taskId string) (*AiTaskStatusRes, error) {
	var res AiTaskStatusRes
	if err := c.getJSON(ctx, fmt.Sprintf("/api/task/status?task_id=%s", taskId), &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ===================== 内部 HTTP helpers =====================

func (c *Client) postJSON(ctx context.Context, path string, body, out interface{}) error {
	b, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseUrl+path, bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("newRequest: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("status %d: %s", resp.StatusCode, string(data))
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func (c *Client) getJSON(ctx context.Context, path string, out interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseUrl+path, nil)
	if err != nil {
		return fmt.Errorf("newRequest: %w", err)
	}
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("status %d: %s", resp.StatusCode, string(data))
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

// readSseStream 发起 POST 请求，以 SSE 协议逐行解析响应，每帧调用 onData 回调
// 使用 http.DefaultClient（无超时），通过 ctx 控制取消
func (c *Client) readSseStream(ctx context.Context, path string, body interface{}, onData func(data string) error) error {
	b, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseUrl+path, bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("newRequest: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	// 流式请求使用无超时的 DefaultClient（ctx 负责取消）
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("do stream: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("status %d: %s", resp.StatusCode, string(data))
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data: ") {
			continue
		}
		data := strings.TrimPrefix(line, "data: ")
		if data == "[DONE]" {
			break
		}
		if err := onData(data); err != nil {
			return fmt.Errorf("onData: %w", err)
		}
	}
	return scanner.Err()
}
