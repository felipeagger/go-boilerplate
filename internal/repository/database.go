package repository

import (
	"github.com/felipeagger/go-boilerplate/internal/entity"
	"gorm.io/gorm"
)

//Migrate execute auto migrations on database
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.User{})
}
