package service

import (
	"dy/model"
	"fmt"
	"testing"
	"time"
)

func TestQueryListInfo(t *testing.T) {
	_ = model.MySQLInit()
	videoList, _, err := QueryListInfo(time.Unix(1694333179, 0))
	if err != nil {
		t.Errorf("出错了%v\n", err)
	}
	fmt.Println(videoList)
}
