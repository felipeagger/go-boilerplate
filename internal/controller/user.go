package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/felipeagger/go-boilerplate/internal/config"
	"github.com/felipeagger/go-boilerplate/pkg/cache"
	"strconv"
	"time"

	"github.com/felipeagger/go-boilerplate/internal/domain"
	"github.com/felipeagger/go-boilerplate/internal/repository"
	"github.com/felipeagger/go-boilerplate/pkg/utils"
)

func init()  {
	cache.InitCacheClientSvc(config.GetEnv().CacheHost, config.GetEnv().CachePort, config.GetEnv().CachePassword)
}

//CreateUser ...
func CreateUser(ctx context.Context, payload domain.Signup) error {

	userRepository := repository.NewGORMUserRepository()

	layOut := "2006/01/02"
	date, _ := time.Parse(layOut, payload.BirthDate)
	birthDate := date
	password, err := utils.GenerateHashPassword(payload.Password)
	if err != nil {
		return errors.New("error on generate hash of password")
	}

	newUser := domain.User{
		ID:        utils.GeneratedID(),
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  password,
		BirthDate: birthDate,
	}

	return userRepository.Create(ctx, newUser)
}

//SignInUser ...
func SignInUser(ctx context.Context, payload domain.Login) (domain.LoginResponse, error) {

	userRepository := repository.NewGORMUserRepository()

	user := userRepository.GetByEmail(ctx, payload.Email)

	if utils.CheckPasswordHash(payload.Password, user.Password) {
		token, err := utils.GenerateJWT(config.GetEnv().TokenSecret, user.ID)
		if err != nil {
			return domain.LoginResponse{
				Message: "fail on create session!",
			}, errors.New("error on generate token")
		}

		err = cache.GetCacheClient().Set(ctx, fmt.Sprintf("tkn-%s", strconv.Itoa(int(user.ID))), token, time.Hour * 1)
		if err != nil {
			return domain.LoginResponse{
				Message: "fail on create session!",
			}, errors.New("error on set token on cache")
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