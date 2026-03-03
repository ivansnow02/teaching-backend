package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ExamQuestionModel = (*customExamQuestionModel)(nil)

type (
	// ExamQuestionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customExamQuestionModel.
	ExamQuestionModel interface {
		examQuestionModel
		FindListByExamId(ctx context.Context, examId uint64) ([]*ExamQuestion, error)
	}

	customExamQuestionModel struct {
		*defaultExamQuestionModel
	}
)

// NewExamQuestionModel returns a model for the database table.
func NewExamQuestionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ExamQuestionModel {
	return &customExamQuestionModel{
		defaultExamQuestionModel: newExamQuestionModel(conn, c, opts...),
	}
}

func (m *customExamQuestionModel) FindListByExamId(ctx context.Context, examId uint64) ([]*ExamQuestion, error) {
	query := fmt.Sprintf("select %s from %s where `exam_id` = ? order by sort asc", examQuestionRows, m.table)
	var list []*ExamQuestion
	err := m.QueryRowsNoCacheCtx(ctx, &list, query, examId)
	if err != nil {
		return nil, err
	}
	return list, nil
}
