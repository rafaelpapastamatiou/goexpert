package database

import (
	"testing"

	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestSaveUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John Doe", "j@j.com", "123456")
	usersRepository := NewGormUsersRepository(db)

	err = usersRepository.Save(user)

	assert.Nil(t, err)

	var userFound entity.User
	err = db.Where("id = ?", user.ID.String()).First(&userFound).Error

	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotEmpty(t, userFound.Password)
}
