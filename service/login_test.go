package service

import (
	"dy/model"
	"testing"
)

func TestLogin(t *testing.T) {
	err := model.MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
	userInfo1 := UserLogin("刀哥", "123")
	if userInfo1.Token != "gg" || userInfo1.State != true {
		t.Error("login error")
	}

	userInfo2 := UserLogin("daodao", "123")
	if userInfo2.State != false {
		t.Error("login error")
	}

	userInfo3 := UserLogin("刀哥", "1234")
	if userInfo3.State != false {
		t.Error("login error")
	}

}
