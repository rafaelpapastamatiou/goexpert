package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/rafaelpapastamatiou/goexpert/18-uow-unit-of-work/internal/entity"
	"github.com/rafaelpapastamatiou/goexpert/18-uow-unit-of-work/internal/repository"
	"github.com/rafaelpapastamatiou/goexpert/18-uow-unit-of-work/pkg/uow"
)

type AddCourseUsecaseInput struct {
	CourseName        string `json:"course_name"`
	CourseDescription string `json:"course_description"`
	CategoryName      string `json:"category_name"`
}

type AddCourseUsecaseOutput struct {
	CourseID string `json:"course_id"`
}

type AddCourseUsecase struct {
	Uow uow.UnitOfWorkInterface
}

func NewAddCourseUsecase(uow uow.UnitOfWorkInterface) *AddCourseUsecase {
	return &AddCourseUsecase{
		Uow: uow,
	}
}

func (a *AddCourseUsecase) Execute(ctx context.Context, input AddCourseUsecaseInput) error {
	return a.Uow.Do(ctx, func() error {
		tmpRepo, err := a.Uow.GetRepository(ctx, "category")
		if err != nil {
			return err
		}
		categoryRepository := tmpRepo.(repository.CategoryRepositoryInterface)

		tmpRepo, err = a.Uow.GetRepository(ctx, "course")
		if err != nil {
			return err
		}
		courseRepository := tmpRepo.(repository.CourseRepositoryInterface)

		category, err := categoryRepository.FindOrCreate(ctx, input.CategoryName)
		if err != nil {
			return err
		}

		err = courseRepository.CreateCourse(ctx, entity.Course{
			ID:          uuid.New().String(),
			Name:        input.CourseName,
			Description: input.CourseDescription,
			CategoryID:  category.ID,
		})
		if err != nil {
			return err
		}

		return nil
	})
}
