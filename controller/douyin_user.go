package controller

import "dy/service"

type UserInfo struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	User       *User  `json:"user"`        // 用户信息
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

func UserInfoQuery(userId string) *UserInfo {
	userInfo := service.UserInfoQuery(userId)
	if !userInfo.State {
		return &UserInfo{
			StatusCode: -1,
			StatusMsg:  "failure",
			User:       nil,
		}
	}
	return &UserInfo{
		StatusCode: 0,
		StatusMsg:  "success",
		User:       (*User)(userInfo.User),
	}

}
