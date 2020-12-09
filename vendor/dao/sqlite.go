package dao

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB definition
var DB *gorm.DB

func InitSqlite() (err error) {
	dsn := "host=localhost user=postgres password=123456 dbname=sp500 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return
}
