package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"method":  "post",
				"message": "Username and password form fields are required",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"method":   "post",
			"username": username,
			"password": password,
			"message":  "login success",
		})
	})

	port := ":8083"
	println("Server (POST Form) running on http://localhost" + port + "/login")
	println("You can use the form from server 1 (login_form_server.go), if it's running and configured to post here, or use curl/Postman.")
	router.Run(port)
}
