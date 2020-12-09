// +heroku goVersion go1.14
module github.com/hoipo/sp500_straddle

go 1.14

replace routers => ./routers

replace logic => ./logic

replace myutils => ./myutils

require (
	dao v0.0.0
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	golang.org/x/sys v0.0.0-20201126233918-771906719818 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gorm.io/driver/postgres v1.0.5 // indirect
	logic v0.0.0
	models v0.0.0
	routers v0.0.0
)

replace models => ./models

replace dao => ./dao
