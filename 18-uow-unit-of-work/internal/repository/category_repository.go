package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/rafaelpapastamatiou/goexpert/18-uow-unit-of-work/internal/db"
	"github.com/rafaelpapastamatiou/goexpert/18-uow-unit-of-work/internal/entity"
)

type CategoryRepositoryInterface interface {
	FindOrCreate(ctx context.Context, name string) (entity.Category, error)
}

type CategoryRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCategoryRepository(sqlDb *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		DB:      sqlDb,
		Queries: db.New(sqlDb),
	}
}

func (c *CategoryRepository) FindOrCreate(ctx context.Context, name string) (entity.Category, error) {
	category, err := c.Queries.GetCategoryByName(ctx, name)
	if err != nil {
		err = c.Queries.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          uuid.New().String(),
			Name:        name,
			Description: sql.NullString{String: name, Valid: true},
		})

		if err != nil {
			return entity.Category{}, err
		}

		category, err = c.Queries.GetCategoryByName(ctx, name)
		if err != nil {
			return entity.Category{}, err
		}
	}

	return entity.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description.String,
	}, nil
}

var _ CategoryRepositoryInterface = &CategoryRepository{}
