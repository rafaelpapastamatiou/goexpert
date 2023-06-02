package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/application/dto"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/domain/entity"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/domain/repository"
)

type UserHandler struct {
	usersRepository repository.UsersRepository
}

func NewUserHandler(
	repository repository.UsersRepository,
) *UserHandler {
	return &UserHandler{
		repository,
	}
}

// Create user godoc
// @Summary Create User
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserRequestBodyDTO true "User"
// @Success 201
// @Failure 500
// @Router /users [post]
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

// AuthenticateUser godoc
// @Summary Authenticate user
// @Description Authenticate user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.AuthenticateUserRequestBodyDTO true "Credentials"
// @Success 200 {object} dto.AuthenticateUserResponseBodyDTO
// @Failure 500
// @Router /users/auth [post]
func (h *UserHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)

	var credentials dto.AuthenticateUserRequestBodyDTO

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.usersRepository.FindByEmail(credentials.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !user.ValidatePassword(credentials.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, err := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := &dto.AuthenticateUserResponseBodyDTO{
		AccessToken: tokenString,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
