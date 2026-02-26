package models

import "gorm.io/gorm"

// Cart 购物车模型
type Cart struct {
    gorm.Model
    UserID    uint    `json:"user_id"`
    ProductID uint    `json:"product_id"`
    Quantity  uint    `json:"quantity"`
    Product   Product `json:"product" gorm:"foreignKey:ProductID"`
} 