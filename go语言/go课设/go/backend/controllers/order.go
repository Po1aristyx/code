package controllers

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateOrder 创建订单
func CreateOrder(c *gin.Context) {
	type Req struct {
		Address string `json:"address"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "参数错误")
		return
	}
	userID := c.GetUint("userID")
	var carts []models.Cart
	if err := models.DB.Preload("Product").Where("user_id = ?", userID).Find(&carts).Error; err != nil || len(carts) == 0 {
		utils.Fail(c, http.StatusBadRequest, "购物车为空")
		return
	}
	total := 0.0
	var items []models.OrderItem
	for _, cart := range carts {
		item := models.OrderItem{
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			Price:     cart.Product.Price,
		}
		items = append(items, item)
		total += float64(cart.Quantity) * cart.Product.Price
	}
	order := models.Order{
		UserID:      userID,
		TotalAmount: total,
		Address:     req.Address,
		Status:      "pending_payment",
		Items:       items,
	}
	if err := models.DB.Create(&order).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, "创建订单失败")
		return
	}
	// 清空购物车
	models.DB.Where("user_id = ?", userID).Delete(&models.Cart{})
	utils.Success(c, order)
}

// GetOrderList 获取当前用户订单列表
func GetOrderList(c *gin.Context) {
	userID := c.GetUint("userID")
	var orders []models.Order
	if err := models.DB.Preload("Items.Product").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, "获取订单失败")
		return
	}
	utils.Success(c, orders)
}

// GetOrderDetail 获取订单详情
func GetOrderDetail(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := models.DB.Preload("Items.Product").First(&order, id).Error; err != nil {
		utils.Fail(c, http.StatusNotFound, "订单不存在")
		return
	}
	utils.Success(c, order)
}

// UpdateOrderStatus 更新订单状态
func UpdateOrderStatus(c *gin.Context) {
	type Req struct {
		Status string `json:"status"`
	}
	var req Req
	id := c.Param("id")
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "参数错误")
		return
	}
	var order models.Order
	if err := models.DB.First(&order, id).Error; err != nil {
		utils.Fail(c, http.StatusNotFound, "订单不存在")
		return
	}
	order.Status = req.Status
	if err := models.DB.Save(&order).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, "更新失败")
		return
	}
	utils.Success(c, order)
} 