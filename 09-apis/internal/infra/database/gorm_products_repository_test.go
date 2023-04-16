package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/domain/entity"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestSaveProduct(t *testing.T) {
	tx, _ := utils.ConnectToTestDatabase(&entity.Product{})
	defer tx.Rollback()

	product, _ := entity.NewProduct("Product 1", 100.0)
	productsRepository := NewGormProductsRepository(tx)

	err := productsRepository.Save(product)

	assert.Nil(t, err)

	var productFound entity.Product
	err = tx.Where("id = ?", product.ID.String()).First(&productFound).Error

	assert.Nil(t, err)
	assert.NotNil(t, productFound)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestFindAllProducts(t *testing.T) {
	tx, _ := utils.ConnectToTestDatabase(&entity.Product{})
	defer tx.Rollback()

	for i := 0; i < 30; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i+1), rand.Float64()*100.0)
		assert.Nil(t, err)

		tx.Create(product)
	}

	productsRepository := NewGormProductsRepository(tx)

	products, err := productsRepository.FindAll(1, 20, "asc")

	assert.NoError(t, err)
	assert.Len(t, products, 20)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 20", products[19].Name)

	products, err = productsRepository.FindAll(2, 20, "asc")

	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 30", products[9].Name)
}

func TestFindProductById(t *testing.T) {
	tx, _ := utils.ConnectToTestDatabase(&entity.Product{})
	defer tx.Rollback()

	product, err := entity.NewProduct("Product 1", 25.99)

	assert.NoError(t, err)

	tx.Create(product)

	productsRepository := NewGormProductsRepository(tx)

	productFound, err := productsRepository.FindById(product.ID.String())

	assert.NoError(t, err)
	assert.NotEmpty(t, productFound)
	assert.Equal(t, productFound.ID, product.ID)
	assert.Equal(t, productFound.Name, product.Name)
	assert.Equal(t, productFound.Price, product.Price)
}

func TestDeleteProduct(t *testing.T) {
	tx, _ := utils.ConnectToTestDatabase(&entity.Product{})
	defer tx.Rollback()

	product, err := entity.NewProduct("Product 1", 14.99)

	assert.NoError(t, err)

	tx.Create(product)

	productsRepository := NewGormProductsRepository(tx)

	err = productsRepository.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productsRepository.FindById(product.ID.String())
	assert.Error(t, err)
}
