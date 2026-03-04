package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"teaching-backend/application/ai/rpc/agentpb"

	"github.com/google/uuid"

	"google.golang.org/grpc"
)

type mockAgentServer struct {
	agentpb.UnimplementedAgentServiceServer
}

func (s *mockAgentServer) AskQuestionStream(req *agentpb.AskQuestionReq, stream agentpb.AgentService_AskQuestionStreamServer) error {
	log.Printf("[AskQuestionStream] course=%d, user=%d, session=%q, question=%q",
		req.CourseId, req.UserId, req.SessionId, req.Question)

	// 如果没有 session_id，生成一个新的
	sessionId := req.SessionId
	if sessionId == "" {
		sessionId = uuid.NewString()
	}

	// 模拟流式回复：第一帧携带 session_id
	deltas := []string{
		"你好！", "我是", "一个",
		"由 Go 编写的 Mock Agent。\n",
		"关于你问的：「", req.Question, "」\n",
		"这是一个 Mock 回复，用于本地测试 SSE 流。\n",
		"✅ session_id = " + sessionId,
	}

	for i, d := range deltas {
		sid := ""
		if i == 0 {
			sid = sessionId // 仅首帧携带
		}
		err := stream.Send(&agentpb.StreamDelta{
			Delta:     d,
			Finished:  false,
			SessionId: sid,
		})
		if err != nil {
			return err
		}
		time.Sleep(200 * time.Millisecond)
	}

	// 最终帧
	return stream.Send(&agentpb.StreamDelta{
		Delta:    "\n[完毕]",
		Sources:  []string{"MockSource-1", "MockSource-2"},
		Finished: true,
	})
}

func (s *mockAgentServer) GenerateCoursewareStream(req *agentpb.GenerateCoursewareReq, stream agentpb.AgentService_GenerateCoursewareStreamServer) error {
	log.Printf("[GenerateCoursewareStream] chapter=%d, user=%d, session=%q",
		req.ChapterId, req.UserId, req.SessionId)

	sessionId := req.SessionId
	if sessionId == "" {
		sessionId = uuid.NewString()
	}

	content := []string{
		"# 模拟课件标题\n",
		"## 第一章：基础知识\n",
		"这是流式生成的课件内容...\n",
		"1. 知识点 A\n",
		"2. 知识点 B\n",
		"3. 知识点 C\n",
		"✅ session_id = " + sessionId,
	}

	for i, c := range content {
		sid := ""
		if i == 0 {
			sid = sessionId // 仅首帧携带
		}
		err := stream.Send(&agentpb.StreamDelta{
			Delta:     c,
			Finished:  false,
			SessionId: sid,
		})
		if err != nil {
			return err
		}
		time.Sleep(300 * time.Millisecond)
	}

	return stream.Send(&agentpb.StreamDelta{
		Delta:    "\n\n--- 生成结束 ---",
		Finished: true,
	})
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	agentpb.RegisterAgentServiceServer(s, &mockAgentServer{})

	fmt.Printf("Mock Agent Server listening at %s\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
