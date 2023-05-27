package dto

type CreateUserRequestBodyDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponseDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
