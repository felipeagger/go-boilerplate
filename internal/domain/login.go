package domain

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
