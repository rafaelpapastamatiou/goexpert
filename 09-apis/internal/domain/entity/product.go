package entity

import (
	"errors"
	"time"

	"github.com/rafaelpapastamatiou/goexpert/09-apis/pkg/entity"
)

var (
	ErrorIdIsRequired    = errors.New("id is required")
	ErrorNameIsRequired  = errors.New("name is required")
	ErrorPriceIsRequired = errors.New("price is required")

	ErrorIdIsInvalid    = errors.New("invalid id")
	ErrorPriceIsInvalid = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	p := Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := p.Validate()
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (p *Product) Validate() error {
	if p.ID == entity.NilID || p.ID.String() == "" {
		return ErrorIdIsRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrorIdIsInvalid
	}

	if p.Name == "" {
		return ErrorNameIsRequired
	}

	if p.Price == 0.0 {
		return ErrorPriceIsRequired
	}

	if p.Price < 0.0 {
		return ErrorPriceIsInvalid
	}

	return nil
}
