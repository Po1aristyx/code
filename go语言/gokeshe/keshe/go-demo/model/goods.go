package model

import (
	_ "github.com/go-sql-driver/mysql"
)

type Goods struct {
	ID       int     `gorm:"primaryKey"`
	Title    string  `gorm:"not null"`
	Name     string  `gorm:"not null;size:255"` //唯一，不为空
	Price    float64 `gorm:"not null;size:255"`
	Details  string  `gorm:"size:255"`
	ImageURL string  `gorm:"not null;size:255"`
	Stock    int     `gorm:"not null"`
}
