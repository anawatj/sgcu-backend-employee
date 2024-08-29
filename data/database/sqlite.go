package db

import (
	"sgcu-backend-employee/config"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectSqlite(configuration *config.Database) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", configuration.DB+".db")

	if err != nil {
		return nil, err
	}
	return db, nil
}
