package controller

import (
    "github.com/gin-gonic/gin"
)

// Route 对所有URL进行路由
// @param server *gin.Engine: 服务器实例
func Route(server *gin.Engine) {
    server.GET("/sync", sync)
}

// sync 处理同步数据的请求
func sync(c *gin.Context) {
    // todo 参数校验
    // todo 调用相应服务
    // todo 返回正确内容
    c.JSON(501, gin.H{
        "message": "Not implemented",
        "result": nil,
    })
}

