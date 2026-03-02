package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StudyProgressModel = (*customStudyProgressModel)(nil)

type (
	// StudyProgressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStudyProgressModel.
	StudyProgressModel interface {
		studyProgressModel
		FindListByUserIdCourseId(ctx context.Context, userId, courseId uint64) ([]*StudyProgress, error)
	}

	customStudyProgressModel struct {
		*defaultStudyProgressModel
	}
)

// NewStudyProgressModel returns a model for the database table.
func NewStudyProgressModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) StudyProgressModel {
	return &customStudyProgressModel{
		defaultStudyProgressModel: newStudyProgressModel(conn, c, opts...),
	}
}

func (m *customStudyProgressModel) FindListByUserIdCourseId(ctx context.Context, userId, courseId uint64) ([]*StudyProgress, error) {
	var list []*StudyProgress
	query := fmt.Sprintf("select %s from %s where user_id = ? and course_id = ?", studyProgressRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &list, query, userId, courseId)
	return list, err
}
