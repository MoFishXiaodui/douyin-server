package model

import (
	"fmt"
	"sync"

	"gorm.io/gorm"
)

// 记录视频和收藏人之间的对应关系
type TableCollect struct {
	gorm.Model
	VideoID uint `json:"video_id"`
	UserID  uint `json:"user_id"`
}

type CollectDao struct {
}

var (
	collectDaoOnce sync.Once
	collectDao     *CollectDao
)

func NewCollectDao() *CollectDao {
	collectDaoOnce.Do(func() {
		collectDao = &CollectDao{}
	})
	return collectDao
}

// 连接到 TableCollect
func ConnectTableCollect() error {
	err := db.AutoMigrate(&TableCollect{})
	if err != nil {
		panic("failed to migrate to collect table")
	}
	fmt.Println("Successful connect to table_collects!(～￣▽￣)～")
	return err
}
