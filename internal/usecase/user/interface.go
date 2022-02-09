package user

import (
	"context"
	"github.com/felipeagger/go-boilerplate/internal/entity"
)

//Reader interface
type Reader interface {
	Get(ctx context.Context, id int64) (user *entity.User, err error)
	GetByEmail(ctx context.Context, email string) (user *entity.User, err error)
}

//Writer user writer
type Writer interface {
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, user *entity.User) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetUser(id int64) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	CreateUser(title string, author string, pages int, quantity int) error
	UpdateUser(user entity.User) error
	DeleteUser(id int64) error
}
