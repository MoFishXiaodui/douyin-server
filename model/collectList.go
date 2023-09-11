package model

import (
	"errors"
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
func TableCollectInit() error {
	err := db.AutoMigrate(&TableCollect{})
	if err != nil {
		panic("failed to migrate to collect table")
	}
	fmt.Println("Successful connect to table_collects!(～￣▽￣)～")
	return err
}

// 创建一条新的收藏记录
func (dao *CollectDao) InsertNewCollectVideoIdUserId(VideoID uint, UserID uint) error {
	// 1. 检查该视频是否存在
	// -------------待补充--------------------------------
	errVideo := VideoInit()
	if errVideo != nil {
		return errors.New("video_table connect failed")
	}

	_, err := NewVideoDao().QueryVideo(VideoID)
	if err != nil {
		return errors.New("video_id is inexistence")
	}

	// 2. 检查该用户是否存在
	errUser := UserInit()
	if errUser != nil {
		return errors.New("user_table connect failed")
	}

	findUserId := NewUserDaoInstance().QuerywithId(UserID)
	if findUserId == nil {
		return errors.New("user_id is inexistence")
	}

	// 3. 查看该用户是否已经收藏过该视频
	errTableCollect := TableCollectInit()
	if errTableCollect != nil {
		return errors.New("table_collects connect failed")
	}

	record, err := NewCollectDao().QueryCollectWithUserID(UserID)
	if err == nil {
		// - 以收藏则返回提示：不能重复收藏
		for i := 0; i < len(record); i++ {
			if record[i].VideoID == VideoID {
				return errors.New("don't collect the same record repeatly")
			}
		}
		// - 不存在则创建新的收藏记录
		res := db.Create(&TableCollect{
			VideoID: VideoID,
			UserID:  UserID,
		})
		if res.Error != nil {
			return res.Error
		}

		// // 给 User 表对应用户的收藏 +1
		// fmt.Println("test 1")
		// err = UserInit()
		// fmt.Println("test 2")
		// if err != nil {
		// 	return errors.New("初始化用户表失败:" + err.Error())
		// }
		// findUser := &User{}
		// fmt.Println("test 3")
		// res = db.Where("id = ?", UserID).First(&findUser)
		// fmt.Println("test 4")
		// fmt.Println(findUser)
		// if res.Error != nil {
		// 	return errors.New("user can't be not found")
		// }
		// fmt.Println("test 5")
		// findUser.FavoriteCount++
		// fmt.Println(findUser)
		// db.Save(&findUser)

		return res.Error
	}
	return err
}

// 查找收藏信息信息（通过 user_id）
func (*CollectDao) QueryCollectWithUserID(UserID uint) ([]TableCollect, error) {
	videos := []TableCollect{}
	res := db.Where("user_id = ?", UserID).Find(&videos)
	if res.Error != nil {
		return nil, errors.New("collect record can't be not found")
	}
	// fmt.Println(videos)
	return videos, nil
}

// 查找收藏信息信息（通过 user_name）
func (*CollectDao) QueryCollectWithUserName(UserName string) ([]TableCollect, error) {
	// 切换到 user_table 找到 name 对应的 user_id
	UserInit()
	user := User{}
	db.Where("name = ?", UserName).First(&user)
	UserID := user.ID

	// 切换到 collect_table 按 user_id 查找表格中的收藏信息
	return NewCollectDao().QueryCollectWithUserID(UserID)
}

// 直接用删除，再插入新 record 复用其他代码
func (*CollectDao) UpdateCollect(OldUserID uint, OldVideoID uint, NewUserID uint, NewVideoID uint) error {
	// 1. 检查老的 record 是否存在
	// 2. 检查修改后的 NewUserID 和 NewVideoID 是否存在
	// 3. 进行修改操作
	err := NewCollectDao().DeleteCollect(OldUserID, OldVideoID)
	if err != nil {
		return errors.New("not find record you want to change")
	}
	return NewCollectDao().InsertNewCollectVideoIdUserId(NewVideoID, NewUserID)
}

// 软删除
func (*CollectDao) DeleteCollect(UserID uint, VideoId uint) error {
	// 先检查该 record 是否存在
	findRecord := &TableCollect{}
	res := db.Where("video_id = ? and user_id = ?", VideoId, UserID).First(&findRecord)
	fmt.Println(findRecord)
	if res.Error != nil {
		return errors.New("collect record can't be not found")
	}

	// 存在则进行软删除
	res = db.Delete(&findRecord)
	return res.Error
}

// 彻底删除
func (*CollectDao) DeleteDeletedCollect(UserID uint, VideoID uint) error {
	res := db.Raw("DELETE FROM table_collects WHERE `user_id` = ? and `video_id` = ? and `deleted_at` IS NOT NULL;", UserID, VideoID).Scan(&TableCollect{UserID: UserID, VideoID: VideoID})
	return res.Error
}
