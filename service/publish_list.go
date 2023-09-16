package service

import (
	"dy/model"
	"errors"
	"fmt"
	"sync"
)

type PublishListQueryFlow struct {
	UserId    uint
	VideoList VideoListDirect
}

func NewPublishListQueryFlow(userId uint) *PublishListQueryFlow {
	// 不同与Dao层，Dao没有自己的状态，Flow有自己的状态，所以Flow每次都要用新的
	return &PublishListQueryFlow{UserId: userId}
}

func (f *PublishListQueryFlow) prepareInfo() error {
	videos, err := model.NewVideoDao().QueryVideosByAuthorId(f.UserId)
	if err != nil {

	}
	list := make(VideoListDirect, len(videos))

	wg := sync.WaitGroup{}
	wg.Add(len(videos))

	for i, v := range videos {
		list[i].CommentCount = v.CommentCount
		list[i].CoverURL = v.CoverUrl
		list[i].FavoriteCount = v.FavoriteCount
		list[i].ID = v.Id
		list[i].Title = v.Title

		//list[i].IsFavorite = 需要其他dao支持，暂不开发

		//list[i].Author
		go func(idx int, playUrl string) {
			defer wg.Done()
			// 处理用户信息
			user := model.NewUserDaoInstance().QuerywithId(f.UserId)
			var totalFavoritedStr string
			if user.TotalFavorited > 1000 {
				totalFavoritedStr = fmt.Sprintf("%.1fk", (float64(user.TotalFavorited))/float64(10))
			} else {
				totalFavoritedStr = fmt.Sprintf("%d", user.TotalFavorited)
			}
			list[idx].Author = User{
				ID:              user.ID,
				Avatar:          user.Avatar,
				BackgroundImage: user.BackgroundImage,
				FavoriteCount:   user.FavoriteCount,
				FollowCount:     user.FollowerCount,
				FollowerCount:   user.FollowerCount,
				IsFollow:        user.IsFollow,
				Name:            user.UserName,
				Signature:       user.Signature,
				TotalFavorited:  totalFavoritedStr,
				WorkCount:       user.WorkCount,
			}

			//	获取视频链接
			url, err := model.NewMinioDao().GetSignedURL(playUrl)
			if err != nil {
				url = "解析错误"
			}
			list[idx].PlayURL = url
		}(i, v.PlayUrl)
	}

	wg.Wait()
	f.VideoList = list
	return nil
}

func (f *PublishListQueryFlow) Do() (VideoListDirect, error) {
	err := f.prepareInfo()
	if err != nil {
		return nil, errors.New("查询出错")
	}
	return f.VideoList, nil
}
