package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 创建一个默认的 Gin 引擎
	r := gin.Default()

	// 2. 定义 /ping 路由
	// 当 GET 请求访问 /ping 路径时，执行指定的处理函数
	r.GET("/ping", func(c *gin.Context) {
		// 3. 返回 JSON 响应
		// http.StatusOK 是 HTTP 状态码 200
		// gin.H 是 map[string]interface{} 的一个快捷方式，用于构建 JSON 对象
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 4. 启动 HTTP 服务，默认监听在 :8080 端口
	// 你也可以指定其他端口，例如 r.Run(":8082")
	err := r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		panic(err) // 如果启动失败，直接 panic
	}
}
