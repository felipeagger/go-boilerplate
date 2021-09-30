package domain

import "time"

type Signup struct {
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	BirthDate time.Time `json:"birthDate"`
	Password  string    `json:"password" binding:"required"`
}
