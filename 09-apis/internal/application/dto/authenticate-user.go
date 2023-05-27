package dto

type AuthenticateUserRequestBodyDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateUserResponseBodyDTO struct {
	AccessToken string `json:"access_token"`
}
