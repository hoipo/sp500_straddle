package main

import (
	"dao"
	"fmt"
	"logic"
	"models"
	"routers"
)

func main() {
	go logic.Job()
	// connect the sqlite database
	err := dao.InitSqlite()
	if err != nil {
		fmt.Printf("init sqlite failed, err:%v\n", err)
		return
	}
	// models binding
	dao.DB.AutoMigrate(&models.FutureData{}, &models.HkETFData{}, &models.HkStockData{}, &models.LofData{})
	// 运行服务
	r := routers.SetupRouters()
	if err := r.Run(); err != nil {
		panic(err)
	}

}
