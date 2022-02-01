package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"github.com/bwmarrin/snowflake"
)

var snowflk *snowflake.Node

func init() {

	var err error
	snowflk, err = snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
	}

}

func GeneratedID() int64 {
	return snowflk.Generate().Int64()
}

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}