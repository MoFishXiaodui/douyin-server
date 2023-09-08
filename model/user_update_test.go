package model

import "testing"

func TestUpdate(t *testing.T) {
	err := MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
	test_user := &User{Name: "daoge", WorkCount: 10}
	real1 := NewUserDaoInstance().Update(1, test_user)
	real2 := NewUserDaoInstance().Update(11, test_user)

	test_user2 := &User{WorkCount: 7}
	real3 := NewUserDaoInstance().Update(1, test_user2)
	if real1 != Success {
		t.Errorf("Expected %v do not match actual %v", Success, real1)
	}
	if real2 != Fail {
		t.Errorf("Expected %v do not match actual %v", Fail, real2)
	}
	if real3 != Success {
		t.Errorf("Expected %v do not match actual %v", Success, real3)
	}

	test_user3 := &User{Name: "小刀", WorkCount: 7}
	real4 := NewUserDaoInstance().Update(1, test_user3)
	if real4 != Fail {
		t.Errorf("Expected %v do not match actual %v", Fail, real4)
	}
}
