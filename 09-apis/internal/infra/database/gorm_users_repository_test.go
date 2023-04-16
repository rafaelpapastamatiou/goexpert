package database

import (
	"testing"

	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/domain/entity"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestSaveUser(t *testing.T) {
	tx, _ := utils.ConnectToTestDatabase(&entity.User{})
	defer tx.Rollback()

	user, _ := entity.NewUser("John Doe", "j@j.com", "123456")
	usersRepository := NewGormUsersRepository(tx)

	err := usersRepository.Save(user)

	assert.Nil(t, err)

	var userFound entity.User
	err = tx.Where("id = ?", user.ID.String()).First(&userFound).Error

	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotEmpty(t, userFound.Password)
}

func TestFindById(t *testing.T) {
	tx, _ := utils.ConnectToTestDatabase(&entity.User{})
	defer tx.Rollback()

	user, _ := entity.NewUser("John Doe", "j@j.com", "123456")
	usersRepository := NewGormUsersRepository(tx)

	err := usersRepository.Save(user)

	assert.Nil(t, err)

	userFound, err := usersRepository.FindById(user.ID.String())

	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, userFound.ID.String(), user.ID.String())
	assert.Equal(t, userFound.Name, "John Doe")
	assert.Equal(t, userFound.Email, "j@j.com")
	assert.NotEmpty(t, userFound.Password)

}

func TestFindByEmail(t *testing.T) {
	tx, _ := utils.ConnectToTestDatabase(&entity.User{})
	defer tx.Rollback()

	user, _ := entity.NewUser("John Doe", "j@j.com", "123456")
	usersRepository := NewGormUsersRepository(tx)

	err := usersRepository.Save(user)

	assert.Nil(t, err)

	userFound, err := usersRepository.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Name, "John Doe")
	assert.Equal(t, userFound.Email, "j@j.com")
	assert.NotEmpty(t, userFound.Password)
}
