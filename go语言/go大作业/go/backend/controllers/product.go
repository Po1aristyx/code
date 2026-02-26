package controllers

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetProductList 获取商品列表
func GetProductList(c *gin.Context) {
	var products []models.Product
	if err := models.DB.Find(&products).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, "获取商品列表失败")
		return
	}
	utils.Success(c, products)
}

// GetProductDetail 获取商品详情
func GetProductDetail(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		utils.Fail(c, http.StatusNotFound, "商品不存在")
		return
	}
	utils.Success(c, product)
} 