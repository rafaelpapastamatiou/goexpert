package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/application/dto"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/domain/entity"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/domain/repository"
)

type UserHandler struct {
	usersRepository repository.UsersRepository
}

func NewUserHandler(repository repository.UsersRepository) *UserHandler {
	return &UserHandler{
		repository,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserRequestBodyDTO

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.usersRepository.Save(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
