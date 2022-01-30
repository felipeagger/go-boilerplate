package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//NewMySQLConnection create a new db connection
func NewMySQLConnection(dbHost, dbName, dbUser, dbPass string) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	//db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	//db.SetMaxIdleConns(cfg.MaxIdleConns)
	//db.SetMaxOpenConns(cfg.MaxOpenConns)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
