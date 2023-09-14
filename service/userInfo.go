package service

import (
	"dy/model"
	"errors"
	"strconv"
)

type QueryUserInfoFlow struct {
	UserId         string          `json:"user_id"`
	UserInfoReturn *UserInfoReturn `json:"user"` // 用户信息
}

type UserInfoReturn struct {
	State bool  `json:"state"` // 描述用户状态
	User  *User `json:"user"`  // 用户信息
}

func (f *QueryUserInfoFlow) prepareUserInfo() error {
	intUserId, _ := strconv.Atoi(f.UserId)
	user := model.NewUserDaoInstance().QuerywithId(uint(intUserId))
	if user == nil {
		return errors.New("this user does not exist")
	}

	f.UserInfoReturn.User = &User{}
	f.UserInfoReturn.User.Avatar = user.Avatar
	f.UserInfoReturn.User.BackgroundImage = user.BackgroundImage
	f.UserInfoReturn.User.FavoriteCount = user.FavoriteCount
	f.UserInfoReturn.User.FollowCount = user.FollowCount
	f.UserInfoReturn.User.FollowerCount = user.FollowerCount
	f.UserInfoReturn.User.ID = user.ID
	f.UserInfoReturn.User.IsFollow = user.IsFollow
	f.UserInfoReturn.User.Name = user.UserName
	f.UserInfoReturn.User.Signature = user.Signature
	f.UserInfoReturn.User.TotalFavorited = strconv.FormatInt(user.TotalFavorited, 10)
	f.UserInfoReturn.User.WorkCount = user.WorkCount

	return nil
}

func UserInfoQuery(userId string) *UserInfoReturn {
	return NewQueryUserInfoFlow(userId).Do()
}

func NewQueryUserInfoFlow(userId string) *QueryUserInfoFlow {
	return &QueryUserInfoFlow{UserId: userId}
}

func (f *QueryUserInfoFlow) Do() *UserInfoReturn {
	f.UserInfoReturn = &UserInfoReturn{}

	user := model.NewUserDaoInstance().QuerywithIdAndToken(f.UserId)
	if user == nil {
		f.UserInfoReturn.State = false
		return f.UserInfoReturn
	}

	if err := f.prepareUserInfo(); err != nil {
		f.UserInfoReturn.State = false
		return f.UserInfoReturn
	}

	f.UserInfoReturn.State = true
	return f.UserInfoReturn
}
