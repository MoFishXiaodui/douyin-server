package controller

import (
	"dy/service"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type DouyinFeed struct {
	// 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	NextTime int64 `json:"next_time"`
	// 状态码，0-成功，其他值-失败
	StatusCode int64 `json:"status_code"`
	// 返回状态描述
	StatusMsg string `json:"status_msg"`
	// 视频列表
	VideoList []service.Video `json:"video_list"`
}

func DouyinFeedGet(latestTimeStr, token string) *DouyinFeed {
	res := &DouyinFeed{}
	latestTime, err := strconv.ParseInt(latestTimeStr, 10, 64)
	if latestTime == 0 {
		fmt.Println("zero")
		latestTime = time.Now().Unix()
	}
	fmt.Println("latestTime", latestTime)
	if err != nil {
		res.StatusCode = http.StatusNotAcceptable
		res.StatusMsg = "时间戳格式不正确"
		return res
	}
	videoList, err := service.QueryListInfo(time.Unix(latestTime, 0))
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.StatusMsg = "服务器发生未知错误"
	}
	res.StatusCode = 0
	res.StatusMsg = "success"
	res.VideoList = videoList.List

	res.NextTime = time.Now().Unix() // 尚未开发
	return res
}
