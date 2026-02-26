package main

import (
    "backend/database"
    "backend/routers"
    "backend/models"
)

func main() {
    database.InitMySQL()
    // 自动迁移表结构
    database.DB.AutoMigrate(&models.User{})
    r := routers.SetupRouter()
    r.Run(":8080")
} 