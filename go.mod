module github.com/hoipo/sp500_straddle

go 1.14

replace routers => ./routers

replace logic => ./logic

replace myutils => ./myutils

require (
	dao v0.0.0
	github.com/cosmtrek/air v1.21.2 // indirect
	github.com/creack/pty v1.1.11 // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	golang.org/x/sys v0.0.0-20201126233918-771906719818 // indirect
	gorm.io/driver/postgres v1.0.5 // indirect
	gorm.io/driver/sqlite v1.1.4 // indirect
	gorm.io/gorm v1.20.7 // indirect
	logic v0.0.0
	models v0.0.0
	routers v0.0.0
)

replace models => ./models

replace dao => ./dao
