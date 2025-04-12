package repository

import (
	"context"
	"database/sql"

	"github.com/rafaelpapastamatiou/goexpert/18-uow-unit-of-work/internal/db"
	"github.com/rafaelpapastamatiou/goexpert/18-uow-unit-of-work/internal/entity"
)

type CourseRepositoryInterface interface {
	CreateCourse(ctx context.Context, course entity.Course) error
}

type CourseRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCourseRepository(sqlDb *sql.DB) *CourseRepository {
	return &CourseRepository{
		DB:      sqlDb,
		Queries: db.New(sqlDb),
	}
}

func (c *CourseRepository) CreateCourse(ctx context.Context, course entity.Course) error {
	// Uncomment the following line to simulate an error and test the rollback
	// return errors.New("course already exists")
	return c.Queries.CreateCourse(ctx, db.CreateCourseParams{
		ID:          course.ID,
		Name:        course.Name,
		Description: sql.NullString{String: course.Description, Valid: true},
		CategoryID:  course.CategoryID,
	})
}

var _ CourseRepositoryInterface = &CourseRepository{}
