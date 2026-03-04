// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package resource

import (
	"context"
	"time"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/pkg/xcode"

	"github.com/minio/minio-go/v7"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUploadTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取 OSS/MinIO 直传签名
func NewGetUploadTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUploadTokenLogic {
	return &GetUploadTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUploadTokenLogic) GetUploadToken() (resp *types.GetUploadTokenRes, err error) {
	bucket := l.svcCtx.Config.Minio.Bucket
	// 设置 dir 为 uploads/yyyy-mm-dd/
	dir := "uploads/" + time.Now().Format("2006-01-02") + "/"

	policy := minio.NewPostPolicy()
	policy.SetBucket(bucket)
	policy.SetKeyStartsWith(dir)
	policy.SetExpires(time.Now().Add(time.Hour * 1))

	_, formData, err := l.svcCtx.MinioClient.PresignedPostPolicy(l.ctx, policy)
	if err != nil {
		return nil, xcode.ServerErr
	}

	host := l.svcCtx.Config.Minio.Endpoint
	if !l.svcCtx.Config.Minio.UseSSL {
		host = "http://" + host
	} else {
		host = "https://" + host
	}
	host = host + "/" + bucket

	return &types.GetUploadTokenRes{
		AccessKeyId: formData["x-amz-credential"],
		Policy:      formData["policy"],
		Signature:   formData["x-amz-signature"],
		Algorithm:   formData["x-amz-algorithm"],
		Date:        formData["x-amz-date"],
		Bucket:      bucket,
		Host:        host,
		Dir:         dir,
	}, nil
}
