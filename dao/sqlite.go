package dao

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB definition
var DB *gorm.DB

func InitSqlite() (err error) {
	// postgres://wgmzwecqpmecdp:25ab5f93b46d85fe0a8e4fe9dd293a5651d6632efb616cbd9d13064f66d36ca6@ec2-3-95-124-37.compute-1.amazonaws.com:5432/df4nb53l52dme0
	dsn := "host=ec2-3-95-124-37.compute-1.amazonaws.com user=wgmzwecqpmecdp password=25ab5f93b46d85fe0a8e4fe9dd293a5651d6632efb616cbd9d13064f66d36ca6 dbname=df4nb53l52dme0 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return
}
