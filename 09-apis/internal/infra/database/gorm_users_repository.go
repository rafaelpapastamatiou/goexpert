package database

import (
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/domain/entity"
	"gorm.io/gorm"
)

type GormUsersRepository struct {
	db *gorm.DB
}

func NewGormUsersRepository(db *gorm.DB) *GormUsersRepository {
	return &GormUsersRepository{
		db: db,
	}
}

func (r *GormUsersRepository) Save(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *GormUsersRepository) FindById(id string) (*entity.User, error) {
	var user entity.User

	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *GormUsersRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
