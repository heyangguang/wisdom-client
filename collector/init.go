package collector

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 启动http服务
func InitHttpListen() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})
	return engine
}
