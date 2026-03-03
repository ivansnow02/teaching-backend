package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ExamModel = (*customExamModel)(nil)

type (
	// ExamModel is an interface to be customized, add more methods here,
	// and implement the added methods in customExamModel.
	ExamModel interface {
		examModel
		FindList(ctx context.Context, courseId uint64, page, pageSize int64) ([]*Exam, int64, error)
	}

	customExamModel struct {
		*defaultExamModel
	}
)

// NewExamModel returns a model for the database table.
func NewExamModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ExamModel {
	return &customExamModel{
		defaultExamModel: newExamModel(conn, c, opts...),
	}
}

func (m *customExamModel) FindList(ctx context.Context, courseId uint64, page, pageSize int64) ([]*Exam, int64, error) {
	where := "where deleted_at is null"
	var vars []any
	if courseId > 0 {
		where += " and course_id = ?"
		vars = append(vars, courseId)
	}

	queryCount := fmt.Sprintf("select count(*) from %s %s", m.table, where)
	var total int64
	err := m.QueryRowNoCacheCtx(ctx, &total, queryCount, vars...)
	if err != nil {
		return nil, 0, err
	}

	if total == 0 {
		return nil, 0, nil
	}

	query := fmt.Sprintf("select %s from %s %s order by id desc limit ?, ?", examRows, m.table, where)
	vars = append(vars, (page-1)*pageSize, pageSize)
	var list []*Exam
	err = m.QueryRowsNoCacheCtx(ctx, &list, query, vars...)
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
