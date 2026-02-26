package main

import (
	"gin-demo/go-demo/common"
	"gin-demo/go-demo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		// 检查请求的路径是否为根路径
		if c.Request.URL.Path == "/" {
			// 使用Redirect方法进行重定向
			c.Redirect(302, "/user/login")
			c.Abort()
			return
		}
		c.Next()
	})

	common.InitDB()

	//创建DB实例进行数据库操作
	db := common.GetDB()
	//延迟关闭数据库
	defer db.Close()
	/*r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "pong",
		})
	})
	r.POST("/user/register", controller.Register) //这里就是用户注册
	r.POST("/user/login", controller.Login)*/
	r = router.InitRouter(r)
	r.Run(":8080")
}
