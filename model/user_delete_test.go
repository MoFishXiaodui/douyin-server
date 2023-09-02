package model

import (
	"testing"
)

func TestDelete(t *testing.T) {
	err := MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
	expect1 := NewUserDaoIstance().Delete(1)
	if expect1 != Success {
		t.Errorf("Expected %v do not match actual %v", expect1, Success)
	}

}
