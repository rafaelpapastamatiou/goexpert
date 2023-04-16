package database

import (
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/domain/entity"
	"gorm.io/gorm"
)

type GormProductsRepository struct {
	db *gorm.DB
}

func NewGormProductsRepository(db *gorm.DB) *GormProductsRepository {
	return &GormProductsRepository{
		db: db,
	}
}

func (r *GormProductsRepository) FindById(id string) (*entity.Product, error) {
	var product entity.Product

	err := r.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *GormProductsRepository) FindAll(page int, limit int, sort string) ([]*entity.Product, error) {
	if sort == "" || (sort != "asc" && sort != "desc") {
		sort = "asc"
	}

	var products []*entity.Product

	tx := r.db.Model(&entity.Product{})

	if page != 0 && limit != 0 {
		tx = tx.Limit(limit).Offset((page - 1) * limit)
	}

	err := tx.Order("created_at " + sort).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *GormProductsRepository) Save(product *entity.Product) error {
	return r.db.Save(product).Error
}

func (r *GormProductsRepository) Delete(id string) error {
	product, err := r.FindById(id)
	if err != nil {
		return err
	}

	return r.db.Delete(product).Error
}
