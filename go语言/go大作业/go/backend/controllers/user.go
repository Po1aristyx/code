package controllers

import (
    "backend/models"
    "backend/utils"
    "backend/database"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "time"
    "golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        utils.Error(c, "参数错误")
        return
    }
    // 密码加密
    hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        utils.Error(c, "密码加密失败")
        return
    }
    user.Password = string(hashedPwd)
    if err := database.DB.Create(&user).Error; err != nil {
        utils.Error(c, "注册失败："+err.Error())
        return
    }
    utils.Success(c, "注册成功")
}

func Login(c *gin.Context) {
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.Error(c, "参数错误")
        return
    }
    var user models.User
    if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
        utils.Error(c, "用户名或密码错误")
        return
    }
    // 校验密码哈希
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        utils.Error(c, "用户名或密码错误")
        return
    }
    // 生成JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "username": user.Username,
        "exp": time.Now().Add(time.Hour * 24 * 7).Unix(), // 7天有效
    })
    tokenString, err := token.SignedString([]byte("your_secret_key"))
    if err != nil {
        utils.Error(c, "生成token失败")
        return
    }
    utils.Success(c, gin.H{"token": tokenString, "user": user})
} 