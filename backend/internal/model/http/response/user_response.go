package response

import "github.com/google/uuid"

type UserRegisterResponse struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
