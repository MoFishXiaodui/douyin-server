package model

import "testing"

func TestUpdate(t *testing.T) {
	err := MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
	test_user := &User{Name: "daoge", WorkCount: 10}
	expect1 := NewUserDaoIstance().Update(1, test_user)
	expect2 := NewUserDaoIstance().Update(11, test_user)
	if expect1 != Success {
		t.Errorf("Expected %v do not match actual %v", expect1, Success)
	}
	if expect2 != Fail {
		t.Errorf("Expected %v do not match actual %v", expect2, Fail)
	}
}
