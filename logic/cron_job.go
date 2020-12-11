package logic

import (
	"fmt"

	"models"

	"github.com/robfig/cron/v3"
)

func Job() {
	c := cron.New() //精确到秒
	fmt.Println("Cron start")

	//定时任务 上海时间05:00记录期货价格
	// c.AddFunc("CRON_TZ=America/New_York 00 16 * * *", func() {
	c.AddFunc("CRON_TZ=America/New_York 30 04 * * *", func() {
		futures, err := GetFuture("ES")
		if err != nil {
			panic(err)
		}
		err = models.CreateFuture(&futures)
		if err != nil {
			panic(err)
		}
		fmt.Println(futures)
	})

	//定时任务 上海时间16:00记录 香港ETF 03140 价格
	// c.AddFunc("CRON_TZ=Asia/Shanghai 01 16 * * *", func() {
	c.AddFunc("CRON_TZ=Asia/Shanghai 20 17 * * *", func() {
		stock, err := GetHKStock("03140")
		if err != nil {
			panic(err)
		}
		err = models.CreateHkStock(&stock)
		if err != nil {
			panic(err)
		}
		fmt.Println(stock)
	})

	//定时任务 上海时间15:01记录 161125 价格
	// c.AddFunc("CRON_TZ=Asia/Shanghai 01 15 * * *", func() {
	c.AddFunc("CRON_TZ=Asia/Shanghai 30 17 * * *", func() {
		lof, err := GetLOF("161125")
		if err != nil {
			panic(err)
		}
		err = models.CreateLof(&lof)
		if err != nil {
			panic(err)
		}
		fmt.Println(lof)
	})

	//定时任务 上海时间09:03记录 161125 价格
	// c.AddFunc("CRON_TZ=Asia/Shanghai 30 09 * * *", func() {
	c.AddFunc("CRON_TZ=Asia/Shanghai 35 17 * * *", func() {
		HkETF, err := GetHkETF()
		if err != nil {
			panic(err)
		}
		err = models.CreateHkETF(&HkETF)
		if err != nil {
			panic(err)
		}
		fmt.Println(HkETF)
	})

	c.Start()
	select {}
}
