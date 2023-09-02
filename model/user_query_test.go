package model

import "testing"

func TestQuery(t *testing.T) {
	err := MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
	expect1 := NewUserDaoIstance().QuerywithName("刀哥")
	expect2 := NewUserDaoIstance().QuerywithName("daoge")
	expect3 := NewUserDaoIstance().QuerywithId(1)
	expect4 := NewUserDaoIstance().QuerywithId(11)

	if expect1 != Existence {
		t.Errorf("Expected %v do not match actual %v", expect1, 1)
	}
	if expect2 != Inexistence {
		t.Errorf("Expected %v do not match actual %v", expect2, 0)
	}

	if expect3 != Existence {
		t.Errorf("Expected %v do not match actual %v", expect1, 1)
	}
	if expect4 != Inexistence {
		t.Errorf("Expected %v do not match actual %v", expect2, 0)
	}
}
