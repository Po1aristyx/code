package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的 Gin 引擎
	router := gin.Default()

	// 定义 /ping 路由
	router.GET("/ping", func(c *gin.Context) {
		// 返回 JSON 响应
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 启动 HTTP 服务器，默认监听在 :8080 端口
	router.Run(":8080")
}
