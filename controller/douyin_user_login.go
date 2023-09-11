package controller

import (
	"dy/config"
	"dy/service"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type UserLoginRsp struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     uint   `json:"user_id"`
	Token      string `json:"token"`
}

func UserLoginPost(username, password string) *UserLoginRsp {
	res := &UserLoginRsp{}
	userinfo := service.UserLogin(username, password)
	if userinfo.State == false {
		res.StatusCode = http.StatusForbidden
		res.StatusMsg = "用户名或密码出错"
		return res
	}

	// token 签发
	key := config.GetJWTconfig()
	t := jwt.New(jwt.SigningMethodHS256)
	s, err := t.SignedString(key)
	if err != nil {
		return &UserLoginRsp{
			StatusCode: -1,
			StatusMsg:  "token generating fail",
			UserId:     0,
			Token:      "",
		}
	}

	res.StatusCode = http.StatusOK
	res.StatusMsg = "success"
	res.UserId = userinfo.UserId
	res.Token = s
	return res
}
