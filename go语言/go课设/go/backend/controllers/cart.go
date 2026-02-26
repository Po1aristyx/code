package controllers

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCart 获取当前用户购物车
func GetCart(c *gin.Context) {
	userID := c.GetUint("userID") // 假设已通过中间件获取
	var carts []models.Cart
	if err := models.DB.Preload("Product").Where("user_id = ?", userID).Find(&carts).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, "获取购物车失败")
		return
	}
	utils.Success(c, carts)
}

// AddToCart 添加商品到购物车
func AddToCart(c *gin.Context) {
	type Req struct {
		ProductID uint `json:"product_id"`
		Quantity  uint `json:"quantity"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "参数错误")
		return
	}
	userID := c.GetUint("userID")
	var cart models.Cart
	if err := models.DB.Where("user_id = ? AND product_id = ?", userID, req.ProductID).First(&cart).Error; err == nil {
		cart.Quantity += req.Quantity
		models.DB.Save(&cart)
	} else {
		cart = models.Cart{UserID: userID, ProductID: req.ProductID, Quantity: req.Quantity}
		models.DB.Create(&cart)
	}
	utils.Success(c, cart)
}

// UpdateCart 更新购物车商品数量
func UpdateCart(c *gin.Context) {
	type Req struct {
		ProductID uint `json:"product_id"`
		Quantity  uint `json:"quantity"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "参数错误")
		return
	}
	userID := c.GetUint("userID")
	var cart models.Cart
	if err := models.DB.Where("user_id = ? AND product_id = ?", userID, req.ProductID).First(&cart).Error; err != nil {
		utils.Fail(c, http.StatusNotFound, "购物车项不存在")
		return
	}
	cart.Quantity = req.Quantity
	models.DB.Save(&cart)
	utils.Success(c, cart)
}

// DeleteCartItem 删除购物车商品
func DeleteCartItem(c *gin.Context) {
	type Req struct {
		ProductID uint `json:"product_id"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "参数错误")
		return
	}
	userID := c.GetUint("userID")
	if err := models.DB.Where("user_id = ? AND product_id = ?", userID, req.ProductID).Delete(&models.Cart{}).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, "删除失败")
		return
	}
	utils.Success(c, nil)
} 