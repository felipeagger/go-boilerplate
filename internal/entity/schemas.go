package entity

type Signup struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	BirthDate string `json:"birthDate"`
	Password  string `json:"password" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
