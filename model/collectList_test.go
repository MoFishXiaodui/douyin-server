package model

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectDao_InsertNewCollectVideoIdUserId_1(t *testing.T) {
	// 插入 user_id 在 user_table 中
	err := MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}

	var expected error = nil
	var user_id, video_id uint
	user_id = 3
	video_id = 6
	res := NewCollectDao().InsertNewCollectVideoIdUserId(video_id, user_id)
	assert.Equal(t, expected, res)

	// 插入重复数据
	res = NewCollectDao().InsertNewCollectVideoIdUserId(video_id, user_id)
	assert.Equal(t, errors.New("don't collect the same record repeatly"), res)
	// 再去数据库检查一下
}

func TestCollectDao_InsertNewCollectVideoIdUserId_2(t *testing.T) {
	// 插入 user_id 不在 user_table 中
	err := MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}

	res := NewCollectDao().InsertNewCollectVideoIdUserId(2, 4)
	assert.Equal(t, errors.New("user_id is inexistence"), res)
	// 再去数据库检查一下
}

func TestCollectDao_QueryCollectWithUserID(t *testing.T) {
	err := MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}

	videos, err := NewCollectDao().QueryCollectWithUserID(2)
	if err != nil {
		fmt.Println("the records is not found")
		return
	}

	fmt.Printf("UserID query sucessfully! 共 %v 条结果，如下：\n", len(videos))
	for i := 0; i < len(videos); i++ {
		fmt.Printf("UserID: %v | VideoID: %v\n", videos[i].UserID, videos[i].VideoID)
	}
	// 对比数据库结果观察用 UserID 查到的记录和数据库中是否一致
}
