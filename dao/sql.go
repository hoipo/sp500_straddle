package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB definition
var DB *gorm.DB

func InitSql() (err error) {
	// postgres://wgmzwecqpmecdp:25ab5f93b46d85fe0a8e4fe9dd293a5651d6632efb616cbd9d13064f66d36ca6@ec2-3-95-124-37.compute-1.amazonaws.com:5432/df4nb53l52dme0
	// dsn := "host=ec2-3-95-124-37.compute-1.amazonaws.com user=wgmzwecqpmecdp password=25ab5f93b46d85fe0a8e4fe9dd293a5651d6632efb616cbd9d13064f66d36ca6 dbname=df4nb53l52dme0 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// dsn := "host=localhost user=postgres password=123456 dbname=sp500 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "root:root1234@tcp(119.27.188.244:3306)/SP500?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return
}
