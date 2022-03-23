package util

import (
	"github.com/jinzhu/gorm"
	"go_gin_api/db"
)

func DBConnect() *gorm.DB {
	var database *gorm.DB
	var err error
	database, err = gorm.Open("postgres", db.Config())
	if err != nil {
		panic("Failed to connect database")
	}
	Migrate(database)
	return database
}
