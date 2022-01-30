package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(200)" json:"name"`
	Email     string `gorm:"type:varchar(100);not null;unique_index" json:"email"`
	BirthDate time.Time
	Password  string `gorm:"type:varchar(255)" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
