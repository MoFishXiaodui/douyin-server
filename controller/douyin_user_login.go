package controller

import (
	"dy/service"
	"net/http"
)

type UserLogin struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     uint   `json:"user_id"`
	Token      string `json:"token"`
}

func UserLoginGet(username, password string) *UserLogin {
	res := &UserLogin{}
	userinfo := service.UserLogin(username, password)
	if userinfo.State == false {
		res.StatusCode = http.StatusInternalServerError
		res.StatusMsg = "服务器发生未知错误"
		return res
	}
	res.StatusCode = http.StatusOK
	res.StatusMsg = "success"
	res.UserId = userinfo.UserId
	res.Token = userinfo.Token
	return res
}
