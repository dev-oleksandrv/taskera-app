package request

type UserRegisterRequest struct {
	Username string `binding:"required" json:"username"`
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required" json:"password"`
}

type UserLoginRequest struct {
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,password" json:"password"`
}
