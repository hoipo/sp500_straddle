package main

import (
	"fmt"

	"sp500_straddle/dao"
	"sp500_straddle/logic"
	"sp500_straddle/models"
	"sp500_straddle/routers"
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
	// port := os.Getenv("PORT")
	if err := r.Run(":8888"); err != nil {
		panic(err)
	}

}
