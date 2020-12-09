module github.com/hoipo/sp500_straddle/logic

go 1.14

replace myutils => ../myutils
replace models => ../models

require myutils v0.0.0
require models v0.0.0

replace dao => ../dao

require dao v0.0.0