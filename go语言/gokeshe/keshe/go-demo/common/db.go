package common

import (
	"fmt"
	"gin-demo/go-demo/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	username := "root"
	password := "Tyx@050529"
	database := "go_demo"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	DB = db
	db.AutoMigrate(&model.User{}) //自动生成表结构
	db.AutoMigrate(&model.Goods{})
	return db

}

func GetDB() *gorm.DB {
	return DB
}
