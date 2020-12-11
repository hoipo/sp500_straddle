package routers

import (
	"net/http"
	"strconv"
	"strings"

	"sp500_straddle/logic"
	"sp500_straddle/models"

	"github.com/gin-gonic/gin"
)

// getFuture 处理函数
func getFuture(c *gin.Context) {
	symbol := c.Param("symbol")
	data, err := logic.GetFuture(symbol)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}

// getLof 处理函数
func getLof(c *gin.Context) {
	symbol := c.Param("symbol")
	data, err := logic.GetLOF(symbol)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}

// getHkETF 处理函数
func getHkETF(c *gin.Context) {
	data, err := logic.GetHkETF()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}

// getForex 处理函数
func getForex(c *gin.Context) {
	symbol := strings.Split(c.Param("symbol"), ",")
	data, err := logic.GetForex(symbol)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}

// getHkStock 处理函数
func getHkStock(c *gin.Context) {
	symbol := c.Param("symbol")
	data, err := logic.GetHKStock(symbol)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}

// getFutureHistory 处理函数
func getFutureHistory(c *gin.Context) {
	limit, err := strconv.ParseInt(c.DefaultQuery("limit", "1"), 10, 8)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	futures, err := models.GetFuture(limit)
	c.JSON(http.StatusOK, futures)
}

// getLofHistory 处理函数
func getLofHistory(c *gin.Context) {
	limit, err := strconv.ParseInt(c.DefaultQuery("limit", "1"), 10, 8)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	futures, err := models.GetLof(limit)
	c.JSON(http.StatusOK, futures)
}

// getHkETFHistory 处理函数
func getHkETFHistory(c *gin.Context) {
	limit, err := strconv.ParseInt(c.DefaultQuery("limit", "1"), 10, 8)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	futures, err := models.GetHkETF(limit)
	c.JSON(http.StatusOK, futures)
}

// getHkStockHistory 处理函数
func getHkStockHistory(c *gin.Context) {
	limit, err := strconv.ParseInt(c.DefaultQuery("limit", "1"), 10, 8)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	futures, err := models.GetHkStock(limit)
	c.JSON(http.StatusOK, futures)
}

// SetupRouters 设定路由
func SetupRouters() *gin.Engine {
	r := gin.Default()

	r.GET("/future/:symbol", getFuture)
	r.GET("/future", getFutureHistory)
	r.GET("/lof/:symbol", getLof)
	r.GET("/lof", getLofHistory)
	r.GET("/hketf/03140", getHkETF)
	r.GET("/hketf", getHkETFHistory)
	r.GET("/hkstock/:symbol", getHkStock)
	r.GET("/hkstock", getHkStockHistory)
	r.GET("/forex/:symbol", getForex)

	return r
}
