package router

import (
	"dy/controller"
	"github.com/gin-gonic/gin"
)

// 接收一个HTTP请求
func DouyinUserLogin(c *gin.Context) {
	var data *controller.UserLogin
	data = controller.UserLoginGet(c.Query("user_name"), c.Query("password"))
	c.JSON(200, data)
}
