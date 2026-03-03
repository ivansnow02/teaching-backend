// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package progress

import (
	"context"
	"encoding/json"
	"fmt"

	"teaching-backend/application/applet/api/internal/code"
	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/pkg/kafkatypes"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStudyProgressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新学习进度
func NewUpdateStudyProgressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStudyProgressLogic {
	return &UpdateStudyProgressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStudyProgressLogic) UpdateStudyProgress(req *types.UpdateStudyProgressReq) (resp *types.Empty, err error) {
	uid, ok := l.ctx.Value("userId").(json.Number)
	if !ok {
		return nil, xcode.AccessDenied
	}
	userId, err := uid.Int64()
	if err != nil || userId <= 0 {
		return nil, xcode.AccessDenied
	}

	// 组装消息并推送到 Kafka，由 course-mq 异步消费落盘
	msg := kafkatypes.StudyProgressMsg{
		UserId:     userId,
		CourseId:   req.CourseId,
		ChapterId:  req.ChapterId,
		MaterialId: req.MaterialId,
		Progress:   int32(req.Progress),
	}

	body, err := json.Marshal(msg)
	if err != nil {
		l.Errorf("序列化学习进度消息失败: %v", err)
		return nil, code.UpdateProgressFailed
	}

	// 以 userId:materialId 作为 key，保证同一用户对同一课件的进度消息有序
	key := fmt.Sprintf("%d:%d", userId, req.MaterialId)
	if err := l.svcCtx.StudyProgressPusher.PushWithKey(l.ctx, key, string(body)); err != nil {
		l.Errorf("推送学习进度到 Kafka 失败: %v", err)
		return nil, code.UpdateProgressFailed
	}

	return &types.Empty{}, nil
}
