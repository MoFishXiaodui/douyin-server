package model

import "testing"

func TestMinioInit(t *testing.T) {
	err := MinioInit()
	if err != nil {
		t.Errorf("sth err: %v\n", err)
	}
}
