package model

import (
	"context"
	"database/sql"
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
		UpdateAiStatus(ctx context.Context, id int64, status int64) error
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

func (m *customCourseMaterialModel) UpdateAiStatus(ctx context.Context, id int64, status int64) error {
	teachingCourseCourseMaterialIdKey := fmt.Sprintf("%s%v", cacheTeachingCourseCourseMaterialIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `ai_status` = ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, status, id)
	}, teachingCourseCourseMaterialIdKey)
	return err
}
