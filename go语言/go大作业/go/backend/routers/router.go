package routers

import (
    "github.com/gin-gonic/gin"
    "backend/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/api/user/register", controllers.Register)
    r.POST("/api/user/login", controllers.Login)
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Go后端服务已启动")
    })
    r.GET("/api/products", controllers.GetProductList)
    r.GET("/api/product/:id", controllers.GetProductDetail)
    // 购物车相关
    r.GET("/api/cart", controllers.GetCart)
    r.POST("/api/cart/add", controllers.AddToCart)
    r.POST("/api/cart/update", controllers.UpdateCart)
    r.POST("/api/cart/delete", controllers.DeleteCartItem)
    // 订单相关
    r.POST("/api/order/create", controllers.CreateOrder)
    r.GET("/api/orders", controllers.GetOrderList)
    r.GET("/api/order/:id", controllers.GetOrderDetail)
    r.POST("/api/order/update_status", controllers.UpdateOrderStatus)
    // 其他路由...
    return r
} 