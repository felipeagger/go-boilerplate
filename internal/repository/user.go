package repository

import (
	"context"
	"github.com/felipeagger/go-boilerplate/internal/entity"
	"github.com/felipeagger/go-boilerplate/pkg/trace"
	"gorm.io/gorm"
)

//UserGORMRepo ...
type UserGORMRepo struct {
	DB *gorm.DB
}

//NewGORMUserRepository ...
func NewGORMUserRepository(dbInstance *gorm.DB) *UserGORMRepo {
	return &UserGORMRepo{
		DB: dbInstance,
	}
}

//Get return a user
func (u *UserGORMRepo) Get(ctx context.Context, id int64) (user *entity.User, err error) {
	ctx, span := trace.NewSpan(ctx, "User.Get")
	defer span.End()

	tx := u.DB.First(&user, "id = ?", id)
	return user, tx.Error
}

//GetByEmail return a user filtered by email
func (u *UserGORMRepo) GetByEmail(ctx context.Context, email string) (user *entity.User, err error) {
	ctx, span := trace.NewSpan(ctx, "User.GetByEmail")
	defer span.End()

	tx := u.DB.First(&user, "email = ?", email)
	return user, tx.Error
}

// Create a user
func (u *UserGORMRepo) Create(ctx context.Context, user *entity.User) error {
	ctx, span := trace.NewSpan(ctx, "User.Create")
	defer span.End()

	result := u.DB.Create(&user)
	return result.Error
}

// Update a user
func (u *UserGORMRepo) Update(ctx context.Context, user *entity.User) error {
	ctx, span := trace.NewSpan(ctx, "User.Update")
	defer span.End()

	result := u.DB.Save(&user)
	return result.Error
}

//Delete a user
func (u *UserGORMRepo) Delete(ctx context.Context, user *entity.User) error {
	ctx, span := trace.NewSpan(ctx, "User.Delete")
	defer span.End()

	result := u.DB.Delete(&user)
	return result.Error
}
