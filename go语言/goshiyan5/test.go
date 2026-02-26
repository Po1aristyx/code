package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	//restful api /article
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	r.GET("/articles", returnAllArticles)
	r.GET("/article/:id", returnSingleArticle)
	r.POST("/article", createNewArticle)
	r.DELETE("/article/:id", deleteArticle)
	r.PUT("/article/:id", updateArticle)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 加载静态页面
	r.LoadHTMLGlob("templates/*")
	// 加载资源文件
	r.Static("/static", "./static")
	//响应一个页面给前端
	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "这是go后台传递来的数据",
		})

	})

	r.GET("/login", func(context *gin.Context) {
		username := context.Query("username")
		password := context.Query("password")
		if len(username) >= 4 && len(username) <= 20 {
			context.JSON(200, gin.H{
				"method":   "get",
				"username": username,
				"password": password,
				"message":  "login success",
			})
		} else {
			context.HTML(http.StatusOK, "login.html", nil)
		}
	})

	r.POST("/login", func(c *gin.Context) {

		username := c.PostForm("username")
		password := c.PostForm("password")

		c.JSON(200, gin.H{
			"method":   "post",
			"username": username,
			"password": password,
			"message":  "login success",
		})
	})

	r.GET("/login/:username/:password", func(c *gin.Context) {
		username := c.Param("username")
		password := c.Param("password")
		c.JSON(200, gin.H{
			"method":   "get restful api",
			"username": username,
			"password": password,
			"message":  "login success",
		})
	})

	r.Run("0.0.0.0:8082") // 监听并在 0.0.0.0:8080 上启动服务

}

type Article struct {
	Id      string `json:Id`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func returnAllArticles(c *gin.Context) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	c.JSON(200, Articles)
}

func returnSingleArticle(c *gin.Context) {
	key := c.Param("id")

	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, article := range Articles {
		if article.Id == key {
			c.JSON(200, article)
		}
	}
}

func createNewArticle(c *gin.Context) {

	var article Article
	if err := c.BindJSON(&article); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	c.JSON(200, Articles)

}

func deleteArticle(c *gin.Context) {
	key := c.Param("id")

	// we then need to loop through all our articles
	for index, article := range Articles {
		// if our id path parameter matches one of our
		// articles
		if article.Id == key {
			// updates our Articles array to remove the
			// article
			Articles = append(Articles[:index], Articles[index+1:]...)
			c.JSON(200, Articles)
		}
	}

}

func updateArticle(c *gin.Context) {

	key := c.Param("id")

	var updateArticle Article
	if err := c.BindJSON(&updateArticle); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, article := range Articles {
		if article.Id == key {
			article = updateArticle
		}
	}

	c.JSON(200, updateArticle)
}
