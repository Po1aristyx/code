package common

import (
	"errors"
	"fmt"
	"gin-demo/go-demo/model"
	"github.com/jinzhu/gorm"
)

func CreateProduct(product model.Goods) error {
	db := GetDB()
	// 使用 Create 方法插入新记录
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func GetAllProducts() ([]model.Goods, error) {
	db := GetDB()
	var products []model.Goods
	// 使用 Find 方法来查询所有记录
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func UpdateProductStock(id int, quantity int) error {
	db := GetDB()

	// 首先，确保库存足够
	var product model.Goods
	if err := db.Where("id = ? AND stock >= ?", id, quantity).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("product not found or stock is insufficient")
		}
		return err
	}

	// 更新库存
	product.Stock -= quantity
	if err := db.Save(&product).Error; err != nil {
		return err
	}

	return nil
}
