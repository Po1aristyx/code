package utils

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Response struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, Response{Code: 0, Msg: "success", Data: data})
}

func Error(c *gin.Context, msg string) {
    c.JSON(http.StatusOK, Response{Code: 1, Msg: msg})
}

func Fail(c *gin.Context, code int, msg string) {
    c.JSON(code, Response{Code: 1, Msg: msg})
} 