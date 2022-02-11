package database

import (
	"fmt"
	gormotel "github.com/wei840222/gorm-otel"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

//NewMySQLConnection create a new db connection
func NewMySQLConnection(dbHost, dbName, dbUser, dbPass string) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.Use(gormotel.New())
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	return db, err
}
