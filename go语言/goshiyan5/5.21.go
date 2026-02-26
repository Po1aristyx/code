package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/login", func(c *gin.Context) {
		// (1) Return an HTML login form
		htmlContent := `
<!DOCTYPE html>
<html>
<head>
    <title>Login Page</title>
</head>
<body>
    <h2>Login Form</h2>
    <hr>
    <h3>POST form</h3>
    <form action="http://localhost:8083/login" method="post">
        <label for="username">Username:</label><br>
        <input type="text" id="username" name="username"><br><br>
        <label for="password">Password:</label><br>
        <input type="password" id="password" name="password"><br><br>
        <input type="submit" value="Login (POST to :8083)">
    </form>
</body>
</html>
`
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})

	port := ":8081"
	println("Server (HTML Form) running on http://localhost" + port + "/login")
	router.Run(port)
}
