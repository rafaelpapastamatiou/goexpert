package repository

import "github.com/rafaelpapastamatiou/goexpert/09-apis/internal/domain/entity"

type ProductsRepository interface {
	Save(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]*entity.Product, error)
	FindById(id string) (*entity.Product, error)
}
