package model

import (
	"errors"
	"sync"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Id            uint
	AuthorId      uint
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	Title         string
}

type VideoDao struct { //帮助理解，这世上有一个猴子种类
}

var (
	VideoDaoOnce sync.Once
	videoDao     *VideoDao //有一只猴子，指向一只猴子
)

func NewVideoDao() *VideoDao {
	VideoDaoOnce.Do(func() {
		videoDao = &VideoDao{} //创建了一只猴子，指向有一只猴子实体
	})
	return videoDao
}

func VideoInit() error {
	return db.AutoMigrate(&Video{})
}

func (dao *VideoDao) CreateNewVideo(title string, AuthorId uint, PlayUrl string) error {
	res := db.Create(&Video{
		AuthorId: AuthorId,
		Title:    title,
		PlayUrl:  PlayUrl,
	})
	return res.Error
}

func (dao *VideoDao) InsertNewVideo(Id, AuthorId uint, FavoriteCount, CommentCount int64, PlayUrl, CoverUrl, title string,
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

func (*VideoDao) QueryVideos() ([]Video, error) {
	results := []Video{}
	res := db.Order("created_at desc").Limit(30).Find(&results)
	if res.Error != nil {
		return nil, errors.New("can't not find")
	}
	return results, nil
}

func (*VideoDao) UpdateVideo(Id, AuthorId uint, FavoriteCount, CommentCount int64, PlayUrl, CoverUrl, title string,
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
