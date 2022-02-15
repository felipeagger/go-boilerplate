package repository

import (
	"github.com/felipeagger/go-boilerplate/internal/entity"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//AutoMigrate execute gorm auto migrations on database
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.User{})
}

//DBMigrate execute all migrations internal.repository.migrations with migrate
func DBMigrate(dbInstance *gorm.DB, dbName string) error {

	db, err := dbInstance.DB()
	if err != nil {
		return err
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return errors.Wrap(err, "failed to instantiate mysql driver")
	}

	migrations, err := migrate.NewWithDatabaseInstance("file://internal/repository/migrations", dbName, driver)
	if err != nil {
		return errors.Wrap(err, "failed to create migrate instance")
	}

	err = migrations.Up()
	if err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "failed to apply migrate up")
	}

	return nil
}