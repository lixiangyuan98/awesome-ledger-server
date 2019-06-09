package controller

import (
	"github.com/gin-gonic/gin"
)

// Route 对所有URL进行路由
// @param server *gin.Engine: 服务器实例
func Route(server *gin.Engine) {
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
}

