package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 15.0)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.NotEmpty(t, p.Name)
	assert.NotEmpty(t, p.Price)
	assert.NotEmpty(t, p.CreatedAt)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 15.0, p.Price)
}

func TestNewProductWhenIdIsEmpty(t *testing.T) {
	p := Product{
		Name:      "Product 1",
		Price:     10.0,
		CreatedAt: time.Now(),
	}

	err := p.Validate()

	assert.Equal(t, ErrorIdIsRequired, err)
}

func TestNewProductWhenNameIsEmpty(t *testing.T) {
	p, err := NewProduct("", 15.0)

	assert.Nil(t, p)
	assert.Equal(t, ErrorNameIsRequired, err)
}

func TestNewProductWhenPriceIsEmpty(t *testing.T) {
	p, err := NewProduct("Product 1", 0)

	assert.Nil(t, p)
	assert.Equal(t, ErrorPriceIsRequired, err)
}

func TestNewProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Product 1", -10.0)

	assert.Nil(t, p)
	assert.Equal(t, ErrorPriceIsInvalid, err)
}

func TestValidateProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 10.0)

	assert.Nil(t, err)
	assert.NotNil(t, p)

	err = p.Validate()

	assert.Nil(t, err)
}
