package kafkatypes

// CanalMsg Canal 推送的 Binlog 变更消息结构
type CanalMsg struct {
	Database string           `json:"database"`
	Table    string           `json:"table"`
	Type     string           `json:"type"` // INSERT, UPDATE, DELETE
	Ts       int64            `json:"ts"`
	IsDDL    bool             `json:"isDdl"`
	Data     []map[string]any `json:"data"`
	Old      []map[string]any `json:"old"`
}

// StudyProgressMsg 学习进度消息结构（applet-api 生产 -> course-mq 消费）
type StudyProgressMsg struct {
	UserId     int64 `json:"userId"`
	CourseId   int64 `json:"courseId"`
	ChapterId  int64 `json:"chapterId"`
	MaterialId int64 `json:"materialId"`
	Progress   int32 `json:"progress"`
}

// SubmitExamMsg 交卷消息结构（exam-rpc 生产 -> exam-mq 消费）
type SubmitExamMsg struct {
	RecordId int64              `json:"recordId"`
	Answers  []SubmitAnswerItem `json:"answers"`
}

type SubmitAnswerItem struct {
	QuestionId int64  `json:"questionId"`
	Answer     string `json:"answer"`
}
