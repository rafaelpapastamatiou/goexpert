package database

import (
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/domain/entity"
	"gorm.io/gorm"
)

type GormUsersRepository struct {
	DB *gorm.DB
}

func NewGormUsersRepository(db *gorm.DB) *GormUsersRepository {
	return &GormUsersRepository{
		DB: db,
	}
}

func (r *GormUsersRepository) Save(user *entity.User) error {
	return r.DB.Create(user).Error
}

func (r *GormUsersRepository) FindById(id string) (*entity.User, error) {
	var user entity.User

	err := r.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *GormUsersRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
