package model

import (
	"errors"
	"strconv"
	"sync"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	//ID int64
	UserName        string
	Password        string
	FollowCount     int64 `gorm:"default:0"`
	FollowerCount   int64 `gorm:"default:0"`
	IsFollow        bool  `gorm:"default:false"`
	Avatar          string
	BackgroundImage string
	Signature       string
	TotalFavorited  int64 `gorm:"default:0"`
	WorkCount       int64 `gorm:"default:0"`
	FavoriteCount   int64 `gorm:"default:0"`
	Token           string
}

type UserDao struct {
}

var (
	userDao  *UserDao
	userOnce sync.Once
)

type UserStatus int64

const (
	//Inexistence UserStatus = iota //用户不存在
	//Existence                     //用户已存在
	Success UserStatus = iota
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
	err = db.First(&u, "user_name = ?", u.UserName).Error
	if err == nil {
		err = errors.New("the name has already been registered")
		return 0, err
	}
	if u.Password == "" || len(u.Password) > 32 {
		return 0, errors.New("please enter the correct password")
	}
	err = db.Create(&u).Error
	if err != nil {
		return 0, err
	}
	return u.ID, nil
}

func (*UserDao) QuerywithName(name string) *User {
	user := &User{}
	err := db.First(user, "user_name = ?", name).Error
	if err != nil {
		return nil
	} else {
		return user
	}
}

func (*UserDao) QuerywithId(id uint) *User {
	user := &User{}
	err := db.First(user, "id = ?", id).Error
	if err != nil {
		return nil
	} else {
		return user
	}
}

func (*UserDao) QueryUsers() ([]User, error) {
	results := []User{}
	res := db.Order("created_at desc").Limit(30).Find(&results)
	if res.Error != nil {
		return nil, errors.New("can't not find")
	}
	return results, nil
}

func (*UserDao) QuerywithNameAndPassword(name, password string) *User {
	user := &User{}
	err := db.First(user, "user_name = ? AND password = ?", name, password).Error
	if err != nil {
		return nil
	} else {
		return user
	}
}

func (*UserDao) QuerywithIdAndToken(userId string) *User {
	user := &User{}
	userIdInt, _ := strconv.Atoi(userId)
	err := db.First(user, "id = ?", uint(userIdInt)).Error
	if err != nil {
		return nil
	} else {
		return user
	}
}

func (*UserDao) Update(id uint, u *User) UserStatus {
	user := &User{}
	user_ := &User{} // 如果新用户名已被使用，则指向那个用户名
	newuser := u

	expect := NewUserDaoInstance().QuerywithId(id)
	if expect != nil {
		db.First(user, "id = ?", id)
		if newuser.UserName != "" {
			// 确认新用户名是否已存在于表格中
			result := db.Where("user_name = ?", newuser.UserName).First(user_).Error
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
	err := db.First(user, "id = ?", id).Delete(user).Error
	if err != nil {
		return Fail
	}
	return Success
}
