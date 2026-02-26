package router

import (
	"gin-demo/go-demo/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {

	//必加才可以访问到
	r.LoadHTMLGlob("D:/QQ数据/gokeshe/keshe/go-demo/template/**/*")
	//加载静态资源
	r.Static("/static", "D:\\QQ数据\\gokeshe\\keshe\\go-demo\\static")
	r.Static("/md", "D:\\QQ数据\\gokeshe\\keshe\\go-demo\\static\\md")
	//首页页面
	r.GET("/user/:id/index", controller.ShowProducts) //加入的页面
	r.POST("/buy_product", controller.BuyProduct)
	//注册
	r.POST("/user/register", controller.Register)
	r.GET("/user/register", controller.RegisterPage) //加入的页面（注意是get）
	//登录
	r.POST("/user/login", controller.Login)
	r.GET("/user/login", controller.LoginPage) //加入的页面（注意是get）

	r.GET("/user/:id", controller.GetUserInfo)
	//r.GET("/user/my", controller.UserPage)//

	r.GET("/user/:id/add_product", controller.AddProductGET)
	r.POST("/user/:id/add_product", controller.AddProductPOST)

	r.POST("/user/password", controller.Newpassword)
	r.GET("/user/password", controller.Password)

	r.GET("/user/:id/product/:goodsid", controller.GetGoodsInfo)
	return r
}
