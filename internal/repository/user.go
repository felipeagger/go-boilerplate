package repository

import (
	"context"
	"github.com/felipeagger/go-boilerplate/internal/domain"
	"github.com/felipeagger/go-boilerplate/pkg/trace"
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
	ctx, span := trace.NewSpan(ctx, "User.Get")
	defer span.End()

	u.DB.First(&user, "id = ?", id)
	return
}

//GetByEmail return a user filtered by email
func (u *User) GetByEmail(ctx context.Context, email string) (user domain.User) {
	ctx, span := trace.NewSpan(ctx, "User.GetByEmail")
	defer span.End()

	u.DB.First(&user, "email = ?", email)
	return
}

// Create a user
func (u *User) Create(ctx context.Context, user domain.User) error {
	ctx, span := trace.NewSpan(ctx, "User.Create")
	defer span.End()

	result := u.DB.Create(&user)
	return result.Error
}

// Update a user
func (u *User) Update(ctx context.Context, user domain.User) error {
	ctx, span := trace.NewSpan(ctx, "User.Update")
	defer span.End()

	result := u.DB.Save(&user)
	return result.Error
}

//Delete a user
func (u *User) Delete(ctx context.Context, user domain.User) error {
	ctx, span := trace.NewSpan(ctx, "User.Delete")
	defer span.End()

	result := u.DB.Delete(&user)
	return result.Error
}
