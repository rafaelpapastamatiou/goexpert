//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/rafaelpapastamatiou/goexpert/19-dependency-injection/product"
)

var repositoryDependencySet = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewProductUseCase(db *sql.DB) *product.ProductUsecase {
	wire.Build(
		repositoryDependencySet,
		product.NewProductUsecase,
	)

	return &product.ProductUsecase{}
}
