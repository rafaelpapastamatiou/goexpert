package product

import "database/sql"

type ProductRepositoryInterface interface {
	GetProductByID(id int) (*Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProductByID(id int) (*Product, error) {
	return &Product{
		ID:   id,
		Name: "Sample Product",
	}, nil
}

var _ ProductRepositoryInterface = &ProductRepository{}

// In-memory implementation of ProductRepositoryInterface
type InMemoryProductRepository struct {
	products map[int]*Product
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: make(map[int]*Product),
	}
}

func (r *InMemoryProductRepository) GetProductByID(id int) (*Product, error) {
	product, exists := r.products[id]
	if !exists {
		return nil, nil
	}
	return product, nil
}

var _ ProductRepositoryInterface = &InMemoryProductRepository{}
