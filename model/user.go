package model

import (
	"errors"
	"gorm.io/gorm"
	"sync"
)

type User struct {
	gorm.Model
	Name            string
	Token           string
	FollowCount     int64 `gorm:"default:0"`
	FollowerCount   int64 `gorm:"default:0"`
	IsFollow        bool  `gorm:"default:false"`
	Avatar          string
	BackgroundImage string
	Signature       string
	TotalFavorited  int64 `gorm:"default:0"`
	WorkCount       int64 `gorm:"default:0"`
	FavoriteCount   int64 `gorm:"default:0"`
}

type UserDao struct {
}

var (
	userDao  *UserDao
	userOnce sync.Once
)

type UserStatus int64

const (
	Inexistence UserStatus = iota //用户不存在
	Existence                     //用户已存在
	Success
	Fail
)

func UserInit() error {
	// Migrate the schema
	return db.AutoMigrate(&User{})
}

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) Create(u User) (id uint, err error) {
	err = db.First(&u, "Name = ?", u.Name).Error
	if err == nil {
		err = errors.New("The name has already been registered")
		return 0, err
	}
	err = db.Create(&u).Error
	if err != nil {
		return 0, err
	}
	return u.ID, nil
}

func (*UserDao) QuerywithName(name string) UserStatus {
	user := &User{}
	err := db.First(user, "name = ?", name).Error
	if err != nil {
		return Inexistence
	} else {
		return Existence
	}
}

func (*UserDao) QuerywithId(id uint) UserStatus {
	user := &User{}
	err := db.First(user, "id = ?", id).Error
	if err != nil {
		return Inexistence
	} else {
		return Existence
	}
}

func (*UserDao) Update(id uint, u *User) UserStatus {
	user := &User{}
	user_ := &User{} // 如果新用户名已被使用，则指向那个用户名
	newuser := u

	expect := NewUserDaoInstance().QuerywithId(id)
	if expect == Existence {
		db.First(user, "id = ?", id)
		if newuser.Name != "" {
			// 确认新用户名是否已存在于表格中
			result := db.Where("name = ?", newuser.Name).First(user_).Error
			if result == nil {
				return Fail
			}
		}
		db.Model(user).Updates(newuser)
		db.Save(user)
		return Success
	} else {
		return Fail
	}
}

func (*UserDao) Delete(id uint) UserStatus {
	//expect := NewUserDaoInstance().QuerywithId(id)
	user := &User{}
	//db.Where("ID = ?", id).Delete(user)
	err := db.First(user, "ID = ?", id).Delete(user).Error
	if err != nil {
		return Fail
	}
	return Success
}
