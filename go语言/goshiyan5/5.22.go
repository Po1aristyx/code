package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/login", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")

		if username == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Username and password query parameters are required.",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"method":   "get",
			"username": username,
			"password": password,
			"message":  "login success",
		})
	})

	port := ":8082"
	println("Server (GET Query Params) running. Test with: http://localhost" + port + "/login?username=test&password=pass")
	router.Run(port)
}
