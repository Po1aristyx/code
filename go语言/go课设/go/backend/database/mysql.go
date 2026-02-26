package database

import (
	"backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitMySQL() {
	dsn := "root:Tyx@050529@tcp(127.0.0.1:3306)/shop1?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	models.DB = DB
	DB.AutoMigrate(&models.Product{}, &models.Cart{}, &models.Order{}, &models.OrderItem{})

	// 初始化商品数据
	var count int64
	DB.Model(&models.Product{}).Count(&count)
	if count == 0 {
		DB.Create(&models.Product{Name: "高性能笔记本电脑", Description: "适合开发和游戏的高性能笔记本", Price: 6999, Image: "https://via.placeholder.com/400x400.png/0000FF/FFFFFF?Text=Product1"})
		DB.Create(&models.Product{Name: "机械键盘", Description: "手感极佳的机械键盘", Price: 799, Image: "https://via.placeholder.com/400x400.png/FF0000/FFFFFF?Text=Product2"})
		DB.Create(&models.Product{Name: "27英寸4K显示器", Description: "超清大屏，护眼不闪屏", Price: 2499, Image: "https://via.placeholder.com/400x400.png/00FF00/FFFFFF?Text=Product3"})
		DB.Create(&models.Product{Name: "无线蓝牙耳机", Description: "降噪舒适，续航持久", Price: 499, Image: "https://via.placeholder.com/400x400.png/FFFF00/000000?Text=Product4"})
		DB.Create(&models.Product{Name: "人体工学椅", Description: "久坐不累，办公首选", Price: 1299, Image: "https://via.placeholder.com/400x400.png/00FFFF/FFFFFF?Text=Product5"})
		DB.Create(&models.Product{Name: "智能手环", Description: "健康监测，运动助手", Price: 299, Image: "https://via.placeholder.com/400x400.png/FF00FF/FFFFFF?Text=Product6"})
		DB.Create(&models.Product{Name: "高清网络摄像头", Description: "高清视频，远程会议必备", Price: 399, Image: "https://via.placeholder.com/400x400.png/C0C0C0/000000?Text=Product7"})
		DB.Create(&models.Product{Name: "大容量移动硬盘", Description: "安全存储，便携高速", Price: 899, Image: "https://via.placeholder.com/400x400.png/808080/FFFFFF?Text=Product8"})
	}
}
