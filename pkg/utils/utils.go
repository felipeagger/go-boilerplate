package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"os"
	"time"

	"github.com/bwmarrin/snowflake"
)

var snowflk *snowflake.Node

func init() {

	var err error
	snowflk, err = snowflake.NewNode(GenerateNodeID())
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

func GenerateNodeID() int64 {
	rand.Seed(time.Now().UnixNano())

	name, _ := os.Hostname()
	if name == "" {
		return int64(rand.Intn(1023))
	}

	id  := int64(rand.Intn(100))
	for _, letter := range []byte(name) {

		if id + int64(letter) >= 1023 {
			return id
		} else {
			id = id + int64(letter)
		}
	}

	return id
}
