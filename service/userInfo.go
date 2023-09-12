package service

import (
	"dy/model"
	"errors"
	"strconv"
)

type QueryUserInfoFlow struct {
	UserId         string          `json:"user_id"` // 状态码，0-成功，其他值-失败
	Token          string          `json:"token"`   // 返回状态描述
	UserInfoReturn *UserInfoReturn `json:"user"`    // 用户信息
}

type UserInfoReturn struct {
	State bool  `json:"state"` // 描述用户状态
	User  *User `json:"user"`  // 用户信息
}

type User struct {
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	ID              uint   `json:"id"`               // 用户id
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Name            string `json:"name"`             // 用户名称
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  string `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
}

func (f *QueryUserInfoFlow) prepareUserInfo() error {
	intUserId, _ := strconv.Atoi(f.UserId)
	user := model.NewUserDaoInstance().QuerywithId(uint(intUserId))
	if user == nil {
		return errors.New("this user does not exist")
	}
	if f.Token != user.Token {
		return errors.New("this token is incorrect")
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
	f.UserInfoReturn.User.TotalFavorited = user.TotalFavorited
	f.UserInfoReturn.User.WorkCount = user.WorkCount

	return nil
}

func UserInfoQuery(user_id, token string) *UserInfoReturn {
	return NewQueryUserInfoFlow(user_id, token).Do()
}

func NewQueryUserInfoFlow(user_id, token string) *QueryUserInfoFlow {
	return &QueryUserInfoFlow{UserId: user_id, Token: token}
}

func (f *QueryUserInfoFlow) Do() *UserInfoReturn {
	f.UserInfoReturn = &UserInfoReturn{}

	user := model.NewUserDaoInstance().QuerywithIdAndToken(f.UserId, f.Token)
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
