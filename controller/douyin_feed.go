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
	res.StatusMsg = "success"

	latestTime, err := strconv.ParseInt(latestTimeStr, 10, 64)
	if err != nil {
		res.StatusCode = 0
		res.StatusMsg = "时间戳格式错误，自动转换成当前时间"
		res.NextTime = time.Now().Unix()
	}
	fmt.Println("err resStatusMsg", latestTimeStr, " -- ", latestTime)
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	videoList, nextTime, err := service.QueryListInfo(time.Unix(latestTime, 0))
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.StatusMsg = "服务器发生未知错误"
		res.NextTime = nextTime.Unix()
		return res
	} else {
		res.StatusCode = 0
		res.VideoList = videoList.List
		res.NextTime = nextTime.Unix()
	}

	return res
}
