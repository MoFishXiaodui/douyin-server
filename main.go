package main

import (
	"dy/model"
	"dy/router"
	"github.com/gin-gonic/gin"
)

func main() {
	err := model.MySQLInit()
	if err != nil {
		panic("mysqlINIT wrong" + err.Error())
	}

	r := gin.Default()
	r.GET("/douyin/feed/", router.DouyinFeed)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
