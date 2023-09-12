package controller

import (
	"dy/model"
	"testing"
)

func TestUserQueryLogin(t *testing.T) {
	err := model.MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}

	loginData := UserLoginPost("刀哥", "123")
	expected := &UserLoginRsp{
		StatusCode: 0,
		StatusMsg:  "success",
		Token:      "gg",
	}
	if (loginData.StatusCode != expected.StatusCode) || (loginData.StatusMsg != expected.StatusMsg) || (loginData.Token != expected.Token) {
		t.Errorf("something wrong happened")
	}

	loginData1 := UserLoginPost("刀哥", "456")
	expected1 := &UserLoginRsp{
		StatusCode: -1,
		StatusMsg:  "failure",
		Token:      "failure",
	}
	if (loginData1.StatusCode != expected1.StatusCode) || (loginData1.StatusMsg != expected1.StatusMsg) || (loginData1.Token != expected1.Token) {
		t.Errorf("something wrong happened")
	}
}
