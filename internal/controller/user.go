package controller

import (
	"context"
	"errors"
	"github.com/felipeagger/go-boilerplate/internal/config"
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

//SignInUser ...
func SignInUser(ctx context.Context, payload domain.Login) (domain.LoginResponse, error) {

	userRepository := repository.NewGORMUserRepository()

	user := userRepository.GetByEmail(ctx, payload.Email)

	if user.Password == payload.Password {
		token, err := utils.GenerateJWT(config.GetEnv().TokenSecret, user.ID)
		if err != nil {
			return domain.LoginResponse{
				Message: "error on generate jwt token!",
			}, errors.New("error on generate token")
		}

		return domain.LoginResponse{
			Token:   token,
			Message: "Success",
		}, nil
	}

	return domain.LoginResponse{
		Message: "Email or Password incorrect!",
	}, errors.New("Unauthorized")
}

//UpdateUser ...
func UpdateUser(ctx context.Context, userID int64, payload domain.Signup) error {

	userRepository := repository.NewGORMUserRepository()

	user := userRepository.Get(ctx, userID)

	if user.ID > 0 {
		user.Email = payload.Email
		user.Name = payload.Name
		user.Password = payload.Password

		return userRepository.Update(ctx, user)
	}

	return nil
}