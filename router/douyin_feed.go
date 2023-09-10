package router

import (
	"dy/controller"
	"github.com/gin-gonic/gin"
)

func DouyinFeed(c *gin.Context) {
	var data *controller.DouyinFeed
	// 获取DouyinFeed内容
	data = controller.DouyinFeedGet(c.Query("latest_time"), c.Query("token"))
	c.JSON(200, data)
}
