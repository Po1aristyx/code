package controller

import (
	"encoding/json"
	"gin-demo/go-demo/common"
	"gin-demo/go-demo/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"time"
)

type PersonalInfo struct {
	Username string `json:"Username"`
	Mobile   string `json:"Mobile"`
	Password int    `json:"Password"`
}

func Newpassword(c *gin.Context) {
	// 获取db
	db := common.GetDB()
	c.Header("Content-Type", "application/json")
	// 创建一个临时结构体来存储 JSON 数据
	type TempUser struct {
		Mobile    string `json:"Mobile"`
		Password  string `json:"Password"`
		Npassword string `json:"Npassword"`
	}
	var tempUser TempUser

	// 读取请求体中的 JSON 数据
	err := json.NewDecoder(c.Request.Body).Decode(&tempUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 422,
			"msg":  "请求格式错误",
		})
		return
	}

	// 使用 tempUser 访问新密码
	newPassword := tempUser.Npassword
	mobile := tempUser.Mobile
	password := tempUser.Password
	//reqUser := model.User{}
	//if err := c.ShouldBindJSON(&reqUser); err != nil {
	//	// 如果绑定失败，返回错误信息
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": 422,
	//		"msg":  "请求格式错误",
	//	})
	//	return
	//}
	//
	//newPassword := c.PostForm("npassword")
	var user model.User
	//mobile := reqUser.Mobile
	//password := reqUser.Password
	user.Mobile = mobile
	user.Password = password

	//判断手机号是否存在
	db.Where("mobile = ?", mobile).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户不存在",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "旧密码错误",
		})
		return
	}
	//密码验证
	if len(newPassword) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "新密码必须大于等于6位",
		})
		return
	}
	//然后对密码进行加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "加密错误",
		})
		return
	}
	if err := db.Model(&model.User{}).Where("mobile = ?", mobile).Updates(model.User{Password: string(hashPassword)}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 422,
			"msg":  "更新错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}

// Register 用户注册
func Register(c *gin.Context) {
	// 获取db
	db := common.GetDB()
	var reqUser model.User
	c.Header("Content-Type", "application/json")
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		// 如果绑定失败，返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 422,
			"msg":  "请求格式错误",
		})
		return
	}

	username := reqUser.Username
	mobile := reqUser.Mobile
	password := reqUser.Password

	//开始数据验证（简单验证）
	if len(mobile) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号不为11位",
		})
		return
	}
	//密码验证
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码必须大于6位",
		})
		return
	}
	//如果用户名为空,生成随机字符串
	if len(username) == 0 {
		username = RandomString(10)
	}

	//判断手机号是否存在
	if isMobileExist(db, mobile) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "当前手机号已经注册",
		})
		return
	}
	//然后对密码进行加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 500,
			"msg":  "加密错误",
		})
		return
	}
	newUser := &model.User{
		Username: username,
		Password: string(hashPassword),
		Mobile:   mobile,
	}
	db.Create(newUser)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

func Login(c *gin.Context) {
	db := common.GetDB()
	var reqUser model.User
	c.Header("Content-Type", "application/json")
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		// 如果绑定失败，返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 422,
			"msg":  "请求格式错误",
		})
		return
	}

	mobile := reqUser.Mobile
	password := reqUser.Password

	var user model.User
	//查询数据
	db.Where("mobile = ?", mobile).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户不存在",
		})
		return
	}
	//密码的对比
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": user,
	})
}

// RandomString 生成随机字符串
func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqlstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789_")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]

	}
	return string(result)

}

func isMobileExist(db *gorm.DB, mobile string) bool {
	var user model.User
	db.Where("mobile = ?", mobile).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// 对应的注册页面的方法
func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

// 对应的登录页面的方法
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
func Password(c *gin.Context) {
	c.HTML(http.StatusOK, "password.html", gin.H{})
}

func GetUserInfo(c *gin.Context) {
	// 从 URL 参数中获取需要查询的用户ID
	userID := c.Param("id")

	// 连接到数据库
	db := common.GetDB()

	// 查询指定ID的用户记录
	var user model.User
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		// 如果查询失败，则向客户端返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	// 将查询结果转换成响应格式
	response := gin.H{
		"id":       user.ID,
		"username": user.Username,
		"Mobile":   user.Mobile,
		// 其他字段...
	}

	// 渲染 HTML 页面并将响应数据传递给模板
	c.HTML(http.StatusOK, "my.html", response)
}
