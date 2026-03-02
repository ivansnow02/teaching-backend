package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CourseModel = (*customCourseModel)(nil)

type (
	// CourseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCourseModel.
	CourseModel interface {
		courseModel
		FindByTeacherId(ctx context.Context, teacherId int64, page, size int64) ([]*Course, int64, error)
	}

	customCourseModel struct {
		*defaultCourseModel
	}
)

// NewCourseModel returns a model for the database table.
func NewCourseModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CourseModel {
	return &customCourseModel{
		defaultCourseModel: newCourseModel(conn, c, opts...),
	}
}

func (m *customCourseModel) FindByTeacherId(ctx context.Context, teacherId int64, page, size int64) ([]*Course, int64, error) {
	var (
		list  []*Course
		total int64
		err   error
	)
	// 查询总数
	countSql := fmt.Sprintf("select count(*) from %s where teacher_id = ?", m.table)
	if teacherId == 0 {
		countSql = fmt.Sprintf("select count(*) from %s", m.table)
		err = m.QueryRowNoCacheCtx(ctx, &total, countSql)
	} else {
		err = m.QueryRowNoCacheCtx(ctx, &total, countSql, teacherId)
	}
	if err != nil {
		return nil, 0, err
	}

	if total == 0 {
		return list, 0, nil
	}

	// 分页查询
	offset := (page - 1) * size
	querySql := fmt.Sprintf("select %s from %s where teacher_id = ? order by create_time desc limit ?, ?", courseRows, m.table)
	if teacherId == 0 {
		querySql = fmt.Sprintf("select %s from %s order by create_time desc limit ?, ?", courseRows, m.table)
		err = m.QueryRowsNoCacheCtx(ctx, &list, querySql, offset, size)
	} else {
		err = m.QueryRowsNoCacheCtx(ctx, &list, querySql, teacherId, offset, size)
	}

	return list, total, err
}
