package controller

import (
	"dy/config"
	"dy/service"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
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
		res.StatusCode = -1
		res.StatusMsg = "用户名或密码出错"
		return res
	}

	// token 签发
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId": userinfo.UserId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenStr, err := t.SignedString(config.GetJWTconfig())
	if err != nil {
		_ = fmt.Errorf("token generating fail %v", err)
		return &UserLoginRsp{
			StatusCode: -1,
			StatusMsg:  "token generating fail",
			UserId:     0,
			Token:      "",
		}
	}

	res.StatusCode = 0
	res.StatusMsg = "success"
	res.UserId = userinfo.UserId
	res.Token = tokenStr
	return res
}
