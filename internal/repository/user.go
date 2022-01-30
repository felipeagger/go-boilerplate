package repository

import (
	"context"

	"github.com/felipeagger/go-boilerplate/internal/domain"
	"gorm.io/gorm"
)

//User is a Relational implementation of UserInterface
type User struct {
	DB *gorm.DB
}

//NewGORMUserRepository ...
func NewGORMUserRepository() UserRepository {
	return &User{
		DB: dbInstance,
	}
}

//Get return a user
func (u *User) Get(ctx context.Context, id int64) (user domain.User) {
	u.DB.First(&user, "id = ?", id)
	return
}

// Create a user
func (u *User) Create(ctx context.Context, user domain.User) error {
	result := u.DB.Create(&user)
	return result.Error
}

// Update a user
func (u *User) Update(ctx context.Context, user domain.User) error {
	result := u.DB.Save(&user)
	return result.Error
}

//Delete a user
func (u *User) Delete(ctx context.Context, user domain.User) error {
	result := u.DB.Delete(&user)
	return result.Error
}
