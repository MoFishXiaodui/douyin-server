package model

import "testing"

func TestQuery(t *testing.T) {
	err := MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
	expect1 := NewUserDaoInstance().QuerywithName("刀哥")
	expect2 := NewUserDaoInstance().QuerywithName("daoge")
	expect3 := NewUserDaoInstance().QuerywithId(1)
	expect4 := NewUserDaoInstance().QuerywithId(11)
	expect5 := NewUserDaoInstance().QuerywithNameAndPassword("刀哥", "123")
	expect6 := NewUserDaoInstance().QuerywithNameAndPassword("刀哥", "456")

	if expect1 == nil {
		t.Errorf("Expected %v do not match actual %v", expect1, nil)
	}
	if expect2 != nil {
		t.Errorf("Expected %v do not match actual a User", expect2)
	}

	if expect3 == nil {
		t.Errorf("Expected %v do not match actual %v", expect3, nil)
	}
	if expect4 != nil {
		t.Errorf("Expected %v do not match actual a User", expect4)
	}
	if expect5 == nil {
		t.Errorf("Expected %v do not match actual %v", expect5, nil)
	}
	if expect6 != nil {
		t.Errorf("Expected %v do not match actual a User", expect6)
	}
}
