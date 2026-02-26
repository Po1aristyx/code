package models

import "gorm.io/gorm"

// Order 订单模型
type Order struct {
    gorm.Model `json:"-"`
    ID         uint        `json:"id" gorm:"primaryKey"`
    UserID      uint        `json:"user_id"`
    TotalAmount float64     `json:"total_amount"`
    Address     string      `json:"address"`
    Status      string      `json:"status"`
    Items       []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

// OrderItem 订单项模型
type OrderItem struct {
    gorm.Model
    OrderID   uint    `json:"order_id"`
    ProductID uint    `json:"product_id"`
    Quantity  uint    `json:"quantity"`
    Price     float64 `json:"price"`
    Product   Product `json:"product" gorm:"foreignKey:ProductID"`
} 