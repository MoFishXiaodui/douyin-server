package model

import (
	"testing"
)

//func TestMain(m *testing.M) {
//	os.Exit(m.Run())
//}

func TestMySQLInit(t *testing.T) {
	err := MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
}
