package controller

import (
	"dy/service"
	"encoding/json"
	"strconv"
)

func UnmarshalPublishList(data []byte) (PublishListRsp, error) {
	var r PublishListRsp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PublishListRsp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PublishListRsp struct {
	StatusCode int64           `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string          `json:"status_msg"`  // 返回状态描述
	VideoList  []service.Video `json:"video_list"`  // 用户发布的视频列表
}

func GetDouyinPublishList(userIdStr string) *PublishListRsp {
	rsp := &PublishListRsp{
		StatusCode: 0,
		StatusMsg:  "success",
	}
	userid64, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		rsp.StatusCode = -1
		rsp.StatusMsg = "用户ID解析错误"
		rsp.VideoList = nil
		return rsp
	}
	userId := uint(userid64)
	list, err := service.NewPublishListQueryFlow(userId).Do()
	if err != nil {
		rsp.StatusCode = -1
		rsp.StatusMsg = "查询列表意外出错"
		rsp.VideoList = nil
		return rsp
	}
	rsp.VideoList = list
	return rsp
}
