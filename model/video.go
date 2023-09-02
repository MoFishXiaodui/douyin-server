package model

import (
	"dy/config"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type Video struct {
	gorm.Model
	Id            uint
	AuthorId      uint
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount uint   `json:"favorite_count"`
	CommentCount  uint   `json:"comment_count"`
	Title         string
}

type VideoDao struct {
}

var (
	VideoDaoOnce sync.Once
	videoDao     *VideoDao
)

func NewVideoDao() *VideoDao {
	VideoDaoOnce.Do(func() {
		videoDao = &VideoDao{}
	})
	return videoDao
}

func InitVideo() {
	addr, user, pwd, dbName := config.GetMySQLConfig()
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user, pwd, addr, dbName)
	fmt.Println("准备连接数据库")
	dbTmp, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("fmt failed to connect database")
		panic("failed to connect database")
	}

	err = dbTmp.AutoMigrate(&Video{})
	if err != nil {
		panic("failed to migrate to videos table")
	}
	fmt.Println("数据库连接成功，videos表对接成功")
	db = dbTmp
}

func (dao *VideoDao) InsertNewVideo(Id, AuthorId, FavoriteCount, CommentCount uint, PlayUrl, CoverUrl, title string,
) error {
	err := dao.DeleteDeletedVideo(Id)
	if err != nil {
		return err
	}
	res := db.Create(&Video{
		Id:            Id,
		AuthorId:      AuthorId,
		PlayUrl:       PlayUrl,
		CoverUrl:      CoverUrl,
		FavoriteCount: FavoriteCount,
		CommentCount:  CommentCount,
		Title:         title,
	})
	return res.Error
}

func (*VideoDao) DeleteDeletedVideo(Id uint) error {
	res := db.Raw("DELETE FROM videos WHERE `id` = ? and `deleted_at` IS NOT NULL ", Id).Scan(&Video{Id: Id})
	return res.Error
}

func (*VideoDao) QueryVideo(Id uint) (*Video, error) {
	v := &Video{Id: Id}
	res := db.First(v)
	if res.Error != nil {
		return nil, errors.New("can't not find")
	}
	return v, nil
}

func (*VideoDao) UpdateVideo(Id, AuthorId, FavoriteCount, CommentCount uint, PlayUrl, CoverUrl, title string,
) error {
	video := &Video{Id: Id}
	firstRes := db.First(video)
	if firstRes.Error != nil {
		return firstRes.Error
	}
	video.AuthorId = AuthorId
	video.PlayUrl = PlayUrl
	video.CoverUrl = CoverUrl
	video.FavoriteCount = FavoriteCount
	video.CommentCount = CommentCount
	video.Title = title
	saveRes := db.Save(video)
	return saveRes.Error
}

func (*VideoDao) UpdateVideoId(newId, oldId uint) error {
	res := db.Table("videos").Where("id = ?", oldId).Updates(map[string]interface{}{"id": newId})
	return res.Error
}

func (*VideoDao) DeleteVideo(Id uint) error {
	video := &Video{Id: Id}
	res := db.First(video)
	if res.Error != nil {
		return res.Error
	}
	res = db.Delete(video)
	return res.Error
}
