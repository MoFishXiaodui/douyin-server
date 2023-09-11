package controller

import "dy/service"

type LoginData struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	Token      string `json:"token"`
}

func ControllerUserQueryLogin(name, password string) *LoginData {
	userinfo := service.UserQueryLogin(name, password)
	if userinfo.State != true {
		return &LoginData{
			StatusCode: -1,
			StatusMsg:  "failure",
			Token:      "failure",
		}
	}
	return &LoginData{
		StatusCode: 0,
		StatusMsg:  "success",
		Token:      userinfo.Token,
	}

}
