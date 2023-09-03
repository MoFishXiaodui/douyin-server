package model

import (
	"errors"
	"gorm.io/gorm"
	"sync"
)

type Favorite struct {
	gorm.Model
	Id         uint
	Lover      string
	IsFavorite bool `json:"is_favorite"`
}

type FavoriteDao struct {
}

var (
	FavoriteDaoOnce sync.Once
	favoriteDao     *FavoriteDao
)

func NewFavoriteDao() *FavoriteDao {
	FavoriteDaoOnce.Do(func() {
		favoriteDao = &FavoriteDao{}
	})
	return favoriteDao
}

func InitFavorite() error {
	return db.AutoMigrate(&Favorite{})
}

func (dao *FavoriteDao) InsertNewFavorite(Id uint, Lover string, IsFavorite bool) error {
	err := dao.DeleteDeleteFavorite(Id)
	if err != nil {
		return err
	}
	res := db.Create(&Favorite{
		Id:         Id,
		Lover:      Lover,
		IsFavorite: IsFavorite,
	})
	return res.Error
}

func (*FavoriteDao) DeleteDeleteFavorite(Id uint) error {
	res := db.Raw("DELETE FROM favorite WHERE `id` = ? and `deleted_at` IS NOT NULL ", Id).Scan(&Favorite{Id: Id})
	return res.Error
}

func (*FavoriteDao) QueryFavorite(Id uint) (*Favorite, error) {
	f := &Favorite{Id: Id}
	res := db.First(f)
	if res.Error != nil {
		return nil, errors.New("can't not find")
	}
	return f, nil
}

func (*FavoriteDao) UpdateFavorite(Id uint, Lover string, IsFavorite bool) error {
	favorite := &Favorite{Id: Id}
	firstRes := db.First(favorite)
	if firstRes.Error != nil {
		return firstRes.Error
	}
	favorite.Id = Id
	favorite.Lover = Lover
	favorite.IsFavorite = IsFavorite
	saveRes := db.Save(favorite)
	return saveRes.Error
}

func (*FavoriteDao) UpdateFavoriteId(newId, oldId uint) error {
	res := db.Table("favorites").Where("id = ?", oldId).Updates(map[string]interface{}{"id": newId})
	return res.Error
}

func (*FavoriteDao) DeleteFavorite(Id uint) error {
	favorite := &Favorite{Id: Id}
	res := db.First(favorite)
	if res.Error != nil {
		return res.Error
	}
	res = db.Delete(favorite)
	return res.Error
}
