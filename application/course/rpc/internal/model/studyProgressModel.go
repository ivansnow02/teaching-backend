package model

import (
	"context"
	"database/sql"
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
		Upsert(ctx context.Context, userId int64, courseId int64, chapterId int64, materialId int64, progress int32) error
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

func (m *customStudyProgressModel) Upsert(ctx context.Context, userId int64, courseId int64, chapterId int64, materialId int64, progress int32) error {
	// 清除缓存
	teachingCourseStudyProgressUserIdMaterialIdKey := fmt.Sprintf("%s%v:%v", cacheTeachingCourseStudyProgressUserIdMaterialIdPrefix, userId, materialId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf(
			"insert into %s (`user_id`, `course_id`, `chapter_id`, `material_id`, `progress`) values (?, ?, ?, ?, ?) "+
				"on duplicate key update `progress` = values(`progress`), `chapter_id` = values(`chapter_id`)",
			m.table,
		)
		return conn.ExecCtx(ctx, query, userId, courseId, chapterId, materialId, progress)
	}, teachingCourseStudyProgressUserIdMaterialIdKey)
	return err
}
