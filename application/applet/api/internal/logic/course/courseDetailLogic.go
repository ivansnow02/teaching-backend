// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package course

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/course/rpc/client/course"

	"github.com/zeromicro/go-zero/core/logx"
)

type CourseDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 课程详情
func NewCourseDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CourseDetailLogic {
	return &CourseDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CourseDetailLogic) CourseDetail(req *types.CourseDetailReq) (resp *types.CourseDetailRes, err error) {
	rpcResp, err := l.svcCtx.CourseRPC.CourseDetail(l.ctx, &course.CourseDetailReq{
		Id: req.Id,
	})
	if err != nil {
		l.Errorf("查询课程详情失败: %v", err)
		return nil, err
	}

	chapters := make([]types.ChapterItem, 0, len(rpcResp.Chapters))
	for _, c := range rpcResp.Chapters {
		materials := make([]types.MaterialItem, 0, len(c.Materials))
		for _, m := range c.Materials {
			materials = append(materials, types.MaterialItem{
				Id:        m.Id,
				ChapterId: m.ChapterId,
				Title:     m.Title,
				Type:      int(m.Type),
				Url:       m.Url,
				FileHash:  m.FileHash,
				FileSize:  m.FileSize,
				AiStatus:  int(m.AiStatus),
				Sort:      int(m.Sort),
			})
		}
		chapters = append(chapters, types.ChapterItem{
			Id:        c.Id,
			CourseId:  c.CourseId,
			Title:     c.Title,
			Sort:      int(c.Sort),
			Materials: materials,
		})
	}

	return &types.CourseDetailRes{
		Course: types.CourseItem{
			Id:          rpcResp.Course.Id,
			Title:       rpcResp.Course.Title,
			Cover:       rpcResp.Course.Cover,
			Description: rpcResp.Course.Description,
			TeacherId:   rpcResp.Course.TeacherId,
			Status:      int(rpcResp.Course.Status),
			CreateTime:  rpcResp.Course.CreateTime,
		},
		Chapters: chapters,
	}, nil
}
