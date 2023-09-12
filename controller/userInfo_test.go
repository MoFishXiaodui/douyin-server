package controller

import (
	"dy/model"
	"testing"
)

func TestControllerUserInfoQuery(t *testing.T) {
	err := model.MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
	userInfo1 := ControllerUserInfoQuery("3", "")
	if userInfo1 == nil {
		t.Error("userInfo1 is nil")
		return
	}
	if userInfo1.User.ID != 3 || userInfo1.User.Name != "tomato" {
		t.Error("userInfoQuery error")
	}
}
