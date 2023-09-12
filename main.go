package main

import (
	"dy/middleware"
	"dy/model"
	"dy/router"
	"github.com/gin-gonic/gin"
)

func main() {
	err := model.MySQLInit()
	if err != nil {
		panic("mysqlINIT wrong" + err.Error())
	}
	err = model.MinioInit()
	if err != nil {
		panic("Minio Init wrong" + err.Error())
	}

	r := gin.Default()

	r.POST("/token_analysis/", router.TokenAnalysisRoute) // token 解析测试

	// 基础接口
	r.GET("/douyin/feed/", router.DouyinFeed)                                          // 视频流
	r.POST("/douyin/user/login/", router.DouyinUserLogin)                              // 用户登录
	r.GET("/douyin/user/", middleware.UserAuth, router.DouyinUser)                     // 用户信息
	r.POST("/douyin/publish/action/", middleware.UserAuth, router.DouyinPublishAction) // 投稿

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
