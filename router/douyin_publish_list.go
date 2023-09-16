package router

import (
	"dy/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DouyinPublishList(c *gin.Context) {
	userId := c.Query("user_id")
	rsp := controller.GetDouyinPublishList(userId)
	if rsp.StatusCode != 0 {
		c.JSON(http.StatusInternalServerError, rsp)
	}
	c.JSON(http.StatusOK, rsp)
}
