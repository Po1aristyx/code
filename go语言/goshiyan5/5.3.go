package main

import (
	"net/http"
	"sync" // To make in-memory store a bit safer for concurrent use (basic)

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Article 结构体定义
type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Author  string `json:"author"`
}

// 使用 map 模拟数据库存储，并添加一个互斥锁以实现基本并发安全
var (
	articlesStore = make(map[string]Article)
	storeLock     = sync.RWMutex{}
)

func main() {
	router := gin.Default()

	// API v1 路由组
	v1 := router.Group("/api/v1")
	{
		articleRoutes := v1.Group("/articles")
		{
			articleRoutes.POST("", createArticle)       // POST /api/v1/articles - 创建文章
			articleRoutes.GET("", getAllArticles)       // GET /api/v1/articles - 获取所有文章
			articleRoutes.GET("/:id", getArticleByID)   // GET /api/v1/articles/:id - 获取单篇文章
			articleRoutes.PUT("/:id", updateArticle)    // PUT /api/v1/articles/:id - 更新文章
			articleRoutes.DELETE("/:id", deleteArticle) // DELETE /api/v1/articles/:id - 删除文章
		}
	}

	port := ":8080"
	println("Article API server running on http://localhost" + port)
	router.Run(port)
}

// --- Handler Functions ---

// createArticle 创建新文章
func createArticle(c *gin.Context) {
	var newArticle Article
	if err := c.ShouldBindJSON(&newArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storeLock.Lock()
	defer storeLock.Unlock()

	newArticle.ID = uuid.New().String() // 生成唯一 ID
	articlesStore[newArticle.ID] = newArticle

	c.JSON(http.StatusCreated, newArticle)
}

// getAllArticles 获取所有文章
func getAllArticles(c *gin.Context) {
	storeLock.RLock()
	defer storeLock.RUnlock()

	// 将 map 转换为 slice 以便返回 JSON 数组
	var articlesList []Article
	for _, article := range articlesStore {
		articlesList = append(articlesList, article)
	}

	c.JSON(http.StatusOK, articlesList)
}

// getArticleByID 根据 ID 获取单篇文章
func getArticleByID(c *gin.Context) {
	id := c.Param("id")

	storeLock.RLock()
	defer storeLock.RUnlock()

	article, exists := articlesStore[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	c.JSON(http.StatusOK, article)
}

// updateArticle 更新指定 ID 的文章
func updateArticle(c *gin.Context) {
	id := c.Param("id")
	var updatedArticle Article

	if err := c.ShouldBindJSON(&updatedArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storeLock.Lock()
	defer storeLock.Unlock()

	_, exists := articlesStore[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	updatedArticle.ID = id // 确保 ID 不变
	articlesStore[id] = updatedArticle

	c.JSON(http.StatusOK, updatedArticle)
}

// deleteArticle 删除指定 ID 的文章
func deleteArticle(c *gin.Context) {
	id := c.Param("id")

	storeLock.Lock()
	defer storeLock.Unlock()

	_, exists := articlesStore[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	delete(articlesStore, id)
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}
