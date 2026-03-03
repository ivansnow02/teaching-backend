package svc

import (
	"teaching-backend/application/course/rpc/internal/config"
	"teaching-backend/application/course/rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                config.Config
	CourseModel           model.CourseModel
	CourseChapterModel    model.CourseChapterModel
	CourseMaterialModel   model.CourseMaterialModel
	StudyProgressModel    model.StudyProgressModel
	CourseEnrollmentModel model.CourseEnrollmentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:                c,
		CourseModel:           model.NewCourseModel(conn, c.CacheRedis),
		CourseChapterModel:    model.NewCourseChapterModel(conn, c.CacheRedis),
		CourseMaterialModel:   model.NewCourseMaterialModel(conn, c.CacheRedis),
		StudyProgressModel:    model.NewStudyProgressModel(conn, c.CacheRedis),
		CourseEnrollmentModel: model.NewCourseEnrollmentModel(conn, c.CacheRedis),
	}
}
