package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// (4) 实现你指定的 GET /login/:username/:password 接口
	router.GET("/login/:username/:password", func(c *gin.Context) {
		// 从 URL 路径中获取参数
		username := c.Param("username")
		password := c.Param("password")

		// 按照你指定的格式返回 JSON 响应
		c.JSON(http.StatusOK, gin.H{
			"method":   "get restful api", // 与你规范中的 "get restful api" 一致
			"username": username,
			"password": password, // 按照你的规范，这里返回了密码
			"message":  "login success",
		})
	})

	port := ":8084" // 你可以根据需要更改端口
	println("Server (implementing your specific GET /login/:username/:password) running.")
	println("Test by opening in your browser or using curl:")
	println("http://localhost" + port + "/login/your_test_user/your_test_password")
	router.Run(port)
}
