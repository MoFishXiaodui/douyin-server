package router

import (
	"dy/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DouyinUser(c *gin.Context) {
	tokenUserId, _ := c.Get("UserId") // tokenUserId - float64

	// 确认用户是否冒用
	// 先把query参数转成整数
	userId, errUserId := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if errUserId != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status_code": -1,
			"status_msg":  "unexpected user_id",
		})
	}
	//fmt.Println(reflect.TypeOf(claims["UserId"]))	// float64
	//fmt.Println(reflect.TypeOf(userId))	// int

	if tokenUserId.(float64) == float64(userId) {
	} else {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status_code": -1,
			"status_msg":  "token 校验失败",
			// 冒用他人userID
		})
		return
	}
	userInfo := controller.UserInfoQuery(strconv.FormatUint(userId, 10))
	c.JSON(http.StatusOK, userInfo)
}
