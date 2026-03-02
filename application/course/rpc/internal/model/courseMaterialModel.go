package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CourseMaterialModel = (*customCourseMaterialModel)(nil)

type (
	// CourseMaterialModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCourseMaterialModel.
	CourseMaterialModel interface {
		courseMaterialModel
		FindListByChapterId(ctx context.Context, chapterId uint64) ([]*CourseMaterial, error)
	}

	customCourseMaterialModel struct {
		*defaultCourseMaterialModel
	}
)

// NewCourseMaterialModel returns a model for the database table.
func NewCourseMaterialModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CourseMaterialModel {
	return &customCourseMaterialModel{
		defaultCourseMaterialModel: newCourseMaterialModel(conn, c, opts...),
	}
}

func (m *customCourseMaterialModel) FindListByChapterId(ctx context.Context, chapterId uint64) ([]*CourseMaterial, error) {
	var list []*CourseMaterial
	query := fmt.Sprintf("select %s from %s where chapter_id = ? order by sort asc, create_time asc", courseMaterialRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &list, query, chapterId)
	return list, err
}
