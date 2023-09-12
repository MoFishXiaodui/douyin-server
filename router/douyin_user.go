package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DouyinUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user_id": c.Query("user_id")})
}
