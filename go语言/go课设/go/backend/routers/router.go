package routers

import (
    "github.com/gin-gonic/gin"
    "backend/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    // 用户相关
    r.POST("/api/users", controllers.Register) // 注册
    r.POST("/api/login", controllers.Login)    // 登录
    // 商品相关
    r.GET("/api/products", controllers.GetProductList)         // 商品列表
    r.GET("/api/products/:id", controllers.GetProductDetail)   // 商品详情
    // 购物车相关
    r.GET("/api/cart", controllers.GetCart)                    // 获取购物车
    r.POST("/api/cart", controllers.AddToCart)                 // 添加商品到购物车
    r.PUT("/api/cart", controllers.UpdateCart)                 // 更新购物车商品数量
    r.DELETE("/api/cart", controllers.DeleteCartItem)          // 删除购物车商品
    // 订单相关
    r.POST("/api/orders", controllers.CreateOrder)             // 创建订单
    r.GET("/api/orders", controllers.GetOrderList)             // 订单列表
    r.GET("/api/orders/:id", controllers.GetOrderDetail)       // 订单详情
    r.PUT("/api/orders/:id/status", controllers.UpdateOrderStatus) // 更新订单状态
    // 主页
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Go后端服务已启动")
    })
    return r
} 