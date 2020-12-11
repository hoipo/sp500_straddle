// +heroku goVersion go1.14
module github.com/hoipo/sp500_straddle

go 1.14

replace routers => ./routers

replace logic => ./logic

replace myutils => ./myutils

require (
	dao v0.0.0
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/robfig/cron/v3 v3.0.0 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	golang.org/x/sys v0.0.0-20201126233918-771906719818 // indirect
	golang.org/x/text v0.3.3 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gorm.io/driver/mysql v1.0.3 // indirect
	logic v0.0.0
	models v0.0.0
	routers v0.0.0
)

replace models => ./models

replace dao => ./dao
