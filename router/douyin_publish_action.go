package router

import (
	"dy/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DouyinPublishAction(c *gin.Context) {
	title := c.Request.FormValue("title")
	userId, _ := c.Get("UserId") // userId - float64

	file, header, _ := c.Request.FormFile("data")
	defer file.Close()

	//contentType := header.Header.Get("Content-Type")
	//fmt.Println("contentType", contentType)	// contentType multipart/form-data -》 不能提供分辨文件类型的作用

	err := service.NewVideoUploadFlow().Upload(title, header.Filename, uint(userId.(float64)), file, header.Size)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status_code": -1,
			"status_msg":  "unknown err: " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "uploaded success",
	})
}
