package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ QuestionModel = (*customQuestionModel)(nil)

type (
	// QuestionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQuestionModel.
	QuestionModel interface {
		questionModel
		FindList(ctx context.Context, courseId, teacherId, qType uint64, page, pageSize int64) ([]*Question, int64, error)
	}

	customQuestionModel struct {
		*defaultQuestionModel
	}
)

// NewQuestionModel returns a model for the database table.
func NewQuestionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) QuestionModel {
	return &customQuestionModel{
		defaultQuestionModel: newQuestionModel(conn, c, opts...),
	}
}

func (m *customQuestionModel) FindList(ctx context.Context, courseId, teacherId, qType uint64, page, pageSize int64) ([]*Question, int64, error) {
	where := "where deleted_at is null"
	var vars []any
	if courseId > 0 {
		where += " and course_id = ?"
		vars = append(vars, courseId)
	}
	if teacherId > 0 {
		where += " and teacher_id = ?"
		vars = append(vars, teacherId)
	}
	if qType > 0 {
		where += " and type = ?"
		vars = append(vars, qType)
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

	query := fmt.Sprintf("select %s from %s %s order by id desc limit ?, ?", questionRows, m.table, where)
	vars = append(vars, (page-1)*pageSize, pageSize)
	var list []*Question
	err = m.QueryRowsNoCacheCtx(ctx, &list, query, vars...)
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
