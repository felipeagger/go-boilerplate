package repository

import (
	"fmt"

	"github.com/felipeagger/go-boilerplate/internal/config"
	"github.com/felipeagger/go-boilerplate/internal/domain"
	"github.com/felipeagger/go-boilerplate/pkg/database"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func init() {

	var err error
	dbInstance, err = database.NewMySQLConnection(config.GetEnv().DBHost, config.GetEnv().DBName, config.GetEnv().DBUser, config.GetEnv().DBPass)
	if err != nil {
		fmt.Println("ERROR DATABASE -> Connection:")
		panic(err)
	}

	migrate(dbInstance)
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&domain.User{})
}
