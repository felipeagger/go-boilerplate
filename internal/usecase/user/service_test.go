package user

import (
	"context"
	"fmt"
	"github.com/felipeagger/go-boilerplate/internal/entity"
	"github.com/felipeagger/go-boilerplate/internal/repository"
	"github.com/felipeagger/go-boilerplate/pkg/cache"
	"github.com/felipeagger/go-boilerplate/pkg/database"
	"github.com/felipeagger/go-boilerplate/pkg/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var userService *Service

func init()  {
	cacheClient, _ := cache.NewCacheMemClient()

	db, err := database.NewSQLiteConnection()
	if err != nil {
		fmt.Println(utils.ErrorDatabaseConn)
		panic(err)
	}

	err = repository.AutoMigrate(db)
	if err != nil {
		fmt.Println(utils.ErrorDatabaseMigrate)
		panic(err)
	}

	userRepository := repository.NewGORMUserRepository(db)
	userService = NewService(userRepository, cacheClient)
}

func TestCreateUser(t *testing.T) {

	err := userService.CreateUser(context.TODO(), entity.Signup{
		Name:      "Satoshi",
		Email:     "satoshi@btc.com",
		Password:  "btc@pswd#",
		BirthDate: "1975/12/31",
	})

	assert.Nil(t, err)

}

func TestSignInUser(t *testing.T) {

	err := userService.CreateUser(context.TODO(), entity.Signup{
		Name:      "Satoshi",
		Email:     "satoshi@btc.org",
		Password:  "btc@pswd#",
		BirthDate: "1975/12/31",
	})

	assert.Nil(t, err)

	//User not found
	resp, err := userService.SignInUser(context.TODO(), entity.Login{
		Email:     "non@exist.com",
		Password:  "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrFindUser, resp.Message)

	// wrong_password
	resp, err = userService.SignInUser(context.TODO(), entity.Login{
		Email:     "satoshi@btc.com",
		Password:  "wrong_password",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "Email or Password incorrect!", resp.Message)

	//Success
	resp, err = userService.SignInUser(context.TODO(), entity.Login{
		Email:     "satoshi@btc.org",
		Password:  "btc@pswd#",
	})

	assert.Nil(t, err)
	assert.Equal(t, "Success", resp.Message)
	assert.NotEmpty(t, resp.Token)
}

func TestUpdateUser(t *testing.T) {

	err := userService.CreateUser(context.TODO(), entity.Signup{
		Name:      "Satoshi",
		Email:     "satoshi@btc.coin",
		Password:  "btc@pswd#",
		BirthDate: "1975/12/31",
	})

	assert.Nil(t, err)

	usr, err := userService.repo.GetByEmail(context.TODO(),"satoshi@btc.coin")
	assert.Nil(t, err)

	err = userService.UpdateUser(context.TODO(), usr.ID, entity.Signup{
		Name:      "Satoshi Nkmto",
		Email:     "satoshi@btc.co",
		Password:  "btc@123#",
		BirthDate: "1976/12/31",
	})

	assert.Nil(t, err)

	usr, err = userService.repo.Get(context.TODO(), usr.ID)

	assert.Nil(t, err)
	assert.Equal(t, "Satoshi Nkmto", usr.Name)
	assert.Equal(t, "satoshi@btc.co", usr.Email)
}

func TestDeleteUser(t *testing.T) {

	err := userService.CreateUser(context.TODO(), entity.Signup{
		Name:      "Satoshi",
		Email:     "satoshi@btc.us",
		Password:  "btc@pswd#",
		BirthDate: "1975/12/31",
	})

	assert.Nil(t, err)

	usr, err := userService.repo.GetByEmail(context.TODO(),"satoshi@btc.us")
	assert.Nil(t, err)

	err = userService.DeleteUser(context.TODO(), usr.ID)

	assert.Nil(t, err)

	usr, err = userService.repo.Get(context.TODO(), usr.ID)

	assert.NotNil(t, err)
	assert.Equal(t, "record not found", err.Error())
	assert.Equal(t, int64(0), usr.ID)
}