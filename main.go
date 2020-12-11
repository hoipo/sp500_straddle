package main

import (
	"dao"
	"fmt"
	"logic"
	"models"
	"os"
	"routers"
)

func main() {
	go logic.Job()
	// connect the sqlite database
	err := dao.InitSql()
	if err != nil {
		fmt.Printf("init sqlite failed, err:%v\n", err)
		return
	}
	// models binding
	dao.DB.AutoMigrate(&models.FutureData{}, &models.HkETFData{}, &models.HkStockData{}, &models.LofData{})
	// 运行服务
	r := routers.SetupRouters()
	port := os.Getenv("PORT")
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}

}
