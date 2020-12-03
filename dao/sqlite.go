package dao

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB definition
var DB *gorm.DB

func InitSqlite() (err error) {
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		return err
	}
	return
}
