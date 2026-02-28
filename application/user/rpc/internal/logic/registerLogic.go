package logic

import (
	"context"
	"time"

	"teaching-backend/application/user/rpc/internal/model"
	"teaching-backend/application/user/rpc/internal/svc"
	"teaching-backend/application/user/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册
func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterRes, error) {
	// 插入用户
	ret, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Email:      in.Email,
		Password:   in.Password,
		Nickname:   in.Nickname,
		Avatar:     "",
		Role:       int64(in.Role),
		Status:     0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	if err != nil {
		l.Errorf("插入用户失败: %v", err)
		return nil, xcode.ServerErr
	}

	userId, err := ret.LastInsertId()
	if err != nil {
		l.Errorf("获取用户ID失败: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.RegisterRes{UserId: userId}, nil
}
