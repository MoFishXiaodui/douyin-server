package service

import (
	"dy/model"
	"errors"
	"strings"
	"time"
)

type VideoList struct {
	List []Video `json:"video_list"` // 用户发布的视频列表
}

type Video struct {
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            uint   `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
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

type QueryListInfoFlow struct {
	LastTime time.Time
	list     *VideoList
}

func (f *QueryListInfoFlow) checkParam() error {

	if f.LastTime.Before(time.Time{}) || strings.ContainsAny(
		f.LastTime.String(), "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return errors.New("Invalid lastTime")
	}
	return nil
}

func (f *QueryListInfoFlow) prepareListInfo() error {
	author := User{"213124sadaf", "23123", 1234, 1235, 4521,
		123, false, "wjl", "sunshine", "snoeeq", 456}
	res, err := model.NewVideoDao().QueryVideos()
	if err != nil {
		return errors.New("获取dao层的video数据出错")
	}
	f.list.List = make([]Video, len(res))
	for i, _ := range res {
		f.list.List[i].ID = res[i].Id
		f.list.List[i].Title = res[i].Title
		f.list.List[i].PlayURL = res[i].PlayUrl
		f.list.List[i].CoverURL = res[i].CoverUrl
		f.list.List[i].FavoriteCount = res[i].FavoriteCount
		f.list.List[i].CommentCount = res[i].CommentCount
		f.list.List[i].Author = author
		f.list.List[i].IsFavorite = true
	}
	return nil
}

func QueryListInfo(lastTime time.Time) (*VideoList, error) {
	return NewQueryListInfoFlow(lastTime).Do()
}

func NewQueryListInfoFlow(lastTime time.Time) *QueryListInfoFlow {
	return &QueryListInfoFlow{LastTime: lastTime}
}

func (f *QueryListInfoFlow) Do() (*VideoList, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareListInfo(); err != nil {
		return nil, err
	}
	return f.list, nil
}
