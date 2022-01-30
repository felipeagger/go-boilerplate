package controller

import (
	"context"
	"time"

	"github.com/felipeagger/go-boilerplate/internal/domain"
	"github.com/felipeagger/go-boilerplate/internal/repository"
	"github.com/felipeagger/go-boilerplate/pkg/utils"
)

//CreateUser ...
func CreateUser(ctx context.Context, payload domain.Signup) error {

	userRepository := repository.NewGORMUserRepository()

	layOut := "2006/01/02"
	date, _ := time.Parse(layOut, payload.BirthDate)
	birthDate := date

	newUser := domain.User{
		ID:        utils.GeneratedID(),
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  payload.Password,
		BirthDate: birthDate,
	}

	return userRepository.Create(ctx, newUser)
}
