package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserAnswerModel = (*customUserAnswerModel)(nil)

type (
	// UserAnswerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAnswerModel.
	UserAnswerModel interface {
		userAnswerModel
		FindListByRecordId(ctx context.Context, recordId uint64) ([]*UserAnswer, error)
	}

	customUserAnswerModel struct {
		*defaultUserAnswerModel
	}
)

// NewUserAnswerModel returns a model for the database table.
func NewUserAnswerModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserAnswerModel {
	return &customUserAnswerModel{
		defaultUserAnswerModel: newUserAnswerModel(conn, c, opts...),
	}
}

func (m *customUserAnswerModel) FindListByRecordId(ctx context.Context, recordId uint64) ([]*UserAnswer, error) {
	query := fmt.Sprintf("select %s from %s where `record_id` = ?", userAnswerRows, m.table)
	var list []*UserAnswer
	err := m.QueryRowsNoCacheCtx(ctx, &list, query, recordId)
	if err != nil {
		return nil, err
	}
	return list, nil
}
