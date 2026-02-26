package models

import "gorm.io/gorm"

// Product 商品模型
type Product struct {
    gorm.Model `json:"-"`
    ID          uint    `json:"id" gorm:"primaryKey"`
    Name        string  `json:"name" gorm:"type:varchar(100);not null"`
    Description string  `json:"description" gorm:"type:text"`
    Price       float64 `json:"price" gorm:"type:decimal(10,2);not null"`
    Image       string  `json:"image" gorm:"type:varchar(255)"`
} 