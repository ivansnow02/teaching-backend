package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CourseChapterModel = (*customCourseChapterModel)(nil)

type (
	// CourseChapterModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCourseChapterModel.
	CourseChapterModel interface {
		courseChapterModel
		FindListByCourseId(ctx context.Context, courseId uint64) ([]*CourseChapter, error)
	}

	customCourseChapterModel struct {
		*defaultCourseChapterModel
	}
)

// NewCourseChapterModel returns a model for the database table.
func NewCourseChapterModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CourseChapterModel {
	return &customCourseChapterModel{
		defaultCourseChapterModel: newCourseChapterModel(conn, c, opts...),
	}
}

func (m *customCourseChapterModel) FindListByCourseId(ctx context.Context, courseId uint64) ([]*CourseChapter, error) {
	var list []*CourseChapter
	query := fmt.Sprintf("select %s from %s where course_id = ? order by sort asc, create_time asc", courseChapterRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &list, query, courseId)
	return list, err
}
