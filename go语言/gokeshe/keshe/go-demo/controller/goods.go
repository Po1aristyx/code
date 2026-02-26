package controller

import (
	"gin-demo/go-demo/common"
	"gin-demo/go-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ShowProducts(c *gin.Context) {
	// 连接到数据库
	db := common.GetDB()

	// 查询所有用户记录
	var goods []model.Goods
	if err := db.Find(&goods).Error; err != nil {
		// 如果查询失败，则向客户端返回错误响应
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "商品查询出错",
		})
		return
	}
	// 调用模型层获取所有产品
	products, err := common.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "商品展示出错",
		})
		return
	}
	// 渲染模板并传递产品数据
	c.HTML(http.StatusOK, "index.html", products)
}

func BuyProduct(c *gin.Context) {
	// 从请求中获取产品 ID 和数量
	productIDStr := c.PostForm("product_id")
	quantityStr := c.PostForm("quantity")

	// 将字符串转换为整数
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		// 如果产品 ID 无效，返回错误信息和 HTTP 400 状态码
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的产品ID",
		})
		return
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		// 如果数量无效，返回错误信息和 HTTP 400 状态码
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的商品数量",
		})
		return
	}

	// 更新产品库存
	if err := common.UpdateProductStock(productID, quantity); err != nil {
		// 如果更新库存失败，返回错误信息和 HTTP 500 状态码
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "更新库存出错",
		})
		return
	}

	// 重定向到产品展示页面
	c.Redirect(http.StatusSeeOther, "/user/:id/index")
}

// AddProductGET 显示添加产品的表单
func AddProductGET(c *gin.Context) {
	c.HTML(http.StatusOK, "add_product.html", nil)
}

// AddProductPOST 处理表单提交并创建新产品
func AddProductPOST(c *gin.Context) {
	var goods model.Goods // 绑定表单数据到 models.Goods 结构体
	if err := c.Bind(&goods); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的输入",
		})
		return
	}
	title := c.PostForm("title")
	name := c.PostForm("name")
	price := c.PostForm("price")
	details := c.PostForm("details")
	imageURL := c.PostForm("image_url")
	stock := c.PostForm("stock")

	price64, err := strconv.ParseFloat(price, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的输入",
		})
	}
	stockInt, err := strconv.Atoi(stock)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的输入",
		})
		return
	}
	goods.Title = title
	goods.Name = name
	goods.Price = price64
	goods.Details = details
	goods.ImageURL = imageURL
	goods.Stock = stockInt

	// 调用 common 包中的 CreateProduct 函数创建新产品
	if err := common.CreateProduct(goods); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建商品出错",
		})
		return
	}

	// 重定向到产品展示页面
	c.Redirect(http.StatusSeeOther, "/user/:id/index")
}

func GetGoodsInfo(c *gin.Context) {
	// 从 URL 参数中获取需要查询的商品ID
	goodsID := c.Param("goodsid")

	// 连接到数据库
	db := common.GetDB()

	// 查询指定ID的用户记录
	var goods model.Goods
	if err := db.Where("id = ?", goodsID).First(&goods).Error; err != nil {
		// 如果查询失败，则向客户端返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	// 将查询结果转换成响应格式
	response := gin.H{
		"ID":       goods.ID,
		"ImageURL": goods.ImageURL,
		"name":     goods.Name,
		"price":    goods.Price,
		"stock":    goods.Stock,
		"details":  goods.Details,
	}

	// 渲染 HTML 页面并将响应数据传递给模板
	c.HTML(http.StatusOK, "goods.html", response)
}
