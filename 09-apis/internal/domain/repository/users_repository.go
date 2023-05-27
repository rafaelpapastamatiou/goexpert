package repository

import "github.com/rafaelpapastamatiou/goexpert/09-apis/internal/domain/entity"

type UsersRepository interface {
	FindById(id string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Save(user *entity.User) error
}
