package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/felipeagger/go-boilerplate/internal/config"
	"github.com/felipeagger/go-boilerplate/pkg/cache"
	"strconv"
	"time"

	"github.com/felipeagger/go-boilerplate/internal/entity"
	"github.com/felipeagger/go-boilerplate/pkg/utils"
)

//Service user usecase
type Service struct {
	repo Repository
	cacheSvc cache.Service
}

//NewService create new service
func NewService(repo Repository, cacheSvc cache.Service) *Service {
	return &Service{
		repo: repo,
		cacheSvc: cacheSvc,
	}
}

//CreateUser ...
func (s *Service) CreateUser(ctx context.Context, payload entity.Signup) error {

	newUser, err := entity.NewUser(payload.Name, payload.Email, payload.Password,
		payload.BirthDate)

	if err != nil {
		return err
	}

	return s.repo.Create(ctx, newUser)
}

//SignInUser ...
func (s *Service) SignInUser(ctx context.Context, payload entity.Login) (entity.LoginResponse, error) {

	user, err := s.repo.GetByEmail(ctx, payload.Email)
	if err != nil {
		return entity.LoginResponse{
			Message: entity.ErrFindUser,
		}, err
	}

	if user.ValidatePassword(payload.Password) {

		token, err := utils.GenerateJWT(config.GetEnv().TokenSecret, user.ID)
		if err != nil {
			return entity.LoginResponse{
				Message: entity.ErrCreateSession,
			}, errors.New("error on generate token")
		}

		err = s.cacheSvc.Set(ctx, fmt.Sprintf("tkn-%s", strconv.Itoa(int(user.ID))), token, time.Hour * 1)
		if err != nil {
			return entity.LoginResponse{
				Message: entity.ErrCreateSession,
			}, errors.New("error on set token on cache")
		}

		return entity.LoginResponse{
			Token:   token,
			Message: "Success",
		}, nil
	}

	return entity.LoginResponse{
		Message: "Email or Password incorrect!",
	}, errors.New("unauthorized")
}

//UpdateUser ...
func (s *Service) UpdateUser(ctx context.Context, userID int64, payload entity.Signup) error {

	user, err := s.repo.Get(ctx, userID)
	if err != nil {
		return err
	}

	if user.ID > 0 {
		user.Email = payload.Email
		user.Name = payload.Name

		err = user.UpdatePassword(payload.Password)
		if err != nil {
			return err
		}

		return s.repo.Update(ctx, user)
	}

	return nil
}

//DeleteUser ...
func (s *Service) DeleteUser(ctx context.Context, userID int64) error {

	user, err := s.repo.Get(ctx, userID)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, user)
}