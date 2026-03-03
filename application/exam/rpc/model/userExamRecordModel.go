package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserExamRecordModel = (*customUserExamRecordModel)(nil)

type (
	// UserExamRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserExamRecordModel.
	UserExamRecordModel interface {
		userExamRecordModel
	}

	customUserExamRecordModel struct {
		*defaultUserExamRecordModel
	}
)

// NewUserExamRecordModel returns a model for the database table.
func NewUserExamRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserExamRecordModel {
	return &customUserExamRecordModel{
		defaultUserExamRecordModel: newUserExamRecordModel(conn, c, opts...),
	}
}
