package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//NewSQLiteConnection create a new db connection
func NewSQLiteConnection() (db *gorm.DB, err error) {

	dsn := "file:memdb1?mode=memory&cache=shared"

	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}