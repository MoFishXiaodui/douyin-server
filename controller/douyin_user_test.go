package controller

import (
	"dy/model"
	"fmt"
	"testing"
)

func TestControllerUserInfoQuery(t *testing.T) {
	err := model.MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
	userInfo1 := UserInfoQuery("3")
	if userInfo1 == nil {
		t.Error("userInfo1 is nil")
		return
	}
	if userInfo1.User.ID != 3 || userInfo1.User.Name != "tomato" {
		t.Error("userInfoQuery error")
	}
}

func TestUserInfoQuery2(t *testing.T) {
	err := model.MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
	userInfo := UserInfoQuery("1")
	if userInfo == nil {
		t.Error("userInfo1 is nil")
		return
	}
	if userInfo.User.ID != 1 || userInfo.User.Name != "刀哥" {
		t.Error("userInfoQuery error")
	}
	fmt.Printf("%#v\n", userInfo.User)
}
