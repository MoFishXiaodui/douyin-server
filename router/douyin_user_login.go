package router

import (
	"dy/controller"
	"github.com/gin-gonic/gin"
)

// 接收一个HTTP请求
func DouyinUserLogin(c *gin.Context) {
	data := controller.UserLoginPost(c.Query("username"), c.Query("password"))
	c.JSON(200, data)
}
