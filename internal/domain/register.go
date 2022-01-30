package domain

type Signup struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	BirthDate string `json:"birthDate"`
	Password  string `json:"password" binding:"required"`
}
