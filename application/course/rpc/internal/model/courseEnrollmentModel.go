package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CourseEnrollmentModel = (*customCourseEnrollmentModel)(nil)

type (
	// CourseEnrollmentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCourseEnrollmentModel.
	CourseEnrollmentModel interface {
		courseEnrollmentModel
		// 分页查询学生已选的课程 (status=1)
		FindListByUserId(ctx context.Context, userId uint64, page, size int64) ([]*CourseEnrollment, int64, error)
		// 分页查询课程已选的学生 (status=1)
		FindListByCourseId(ctx context.Context, courseId uint64, page, size int64) ([]*CourseEnrollment, int64, error)
	}

	customCourseEnrollmentModel struct {
		*defaultCourseEnrollmentModel
	}
)

// NewCourseEnrollmentModel returns a model for the database table.
func NewCourseEnrollmentModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CourseEnrollmentModel {
	return &customCourseEnrollmentModel{
		defaultCourseEnrollmentModel: newCourseEnrollmentModel(conn, c, opts...),
	}
}

// FindListByUserId 按学生ID分页查询已选课程（status=1）
func (m *customCourseEnrollmentModel) FindListByUserId(ctx context.Context, userId uint64, page, size int64) ([]*CourseEnrollment, int64, error) {
	var total int64
	countSQL := fmt.Sprintf("select count(*) from %s where `user_id` = ? and `status` = 1", m.table)
	if err := m.QueryRowNoCacheCtx(ctx, &total, countSQL, userId); err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, nil
	}
	offset := (page - 1) * size
	var list []*CourseEnrollment
	querySQL := fmt.Sprintf("select %s from %s where `user_id` = ? and `status` = 1 order by `create_time` desc limit ?, ?", courseEnrollmentRows, m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &list, querySQL, userId, offset, size); err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// FindListByCourseId 按课程ID分页查询已选学生（status=1）
func (m *customCourseEnrollmentModel) FindListByCourseId(ctx context.Context, courseId uint64, page, size int64) ([]*CourseEnrollment, int64, error) {
	var total int64
	countSQL := fmt.Sprintf("select count(*) from %s where `course_id` = ? and `status` = 1", m.table)
	if err := m.QueryRowNoCacheCtx(ctx, &total, countSQL, courseId); err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, nil
	}
	offset := (page - 1) * size
	var list []*CourseEnrollment
	querySQL := fmt.Sprintf("select %s from %s where `course_id` = ? and `status` = 1 order by `create_time` desc limit ?, ?", courseEnrollmentRows, m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &list, querySQL, courseId, offset, size); err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
