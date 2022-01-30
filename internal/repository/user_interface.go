package repository

import (
	"context"

	"github.com/felipeagger/go-boilerplate/internal/domain"
)

type UserRepository interface {
	Get(ctx context.Context, id int64) domain.User
	Create(ctx context.Context, user domain.User) error
	Update(ctx context.Context, user domain.User) error
	Delete(ctx context.Context, user domain.User) error
}
