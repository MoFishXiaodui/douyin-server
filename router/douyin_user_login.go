package router

import (
	"dy/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 接收一个HTTP请求
func DouyinUserLogin(c *gin.Context) {
	data := controller.UserLoginPost(c.Query("username"), c.Query("password"))
	if data.StatusCode == -1 {
		c.JSON(http.StatusForbidden, data)
	} else if data.StatusCode == 0 {
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_msg":  "unknown err",
			"status_code": "-2",
		})
	}
}
