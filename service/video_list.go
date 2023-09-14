package service

import (
	"dy/model"
	"errors"
	"time"
)

type QueryListInfoFlow struct {
	LastTime time.Time
	NextTime time.Time
	list     *VideoList
}

func (f *QueryListInfoFlow) checkParam() error {

	//if f.LastTime.Before(time.Time{}) || strings.ContainsAny(
	//	f.LastTime.String(), "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
	//	return errors.New("Invalid lastTime")
	//}
	return nil
}

func (f *QueryListInfoFlow) prepareListInfo() (error, time.Time) {
	//author := User{"213124sadaf", "23123", 1234, 1235, 4521,
	//	1, false, "wjl", "sunshine", "snoeeq", 456}
	author := User{}
	res, err := model.NewVideoDao().QueryVideos()
	if err != nil {
		return errors.New("获取dao层的video数据出错"), time.Now()
	}
	f.list = &VideoList{}
	f.list.List = make([]Video, len(res))
	if len(res) == 0 {
		return nil, time.Now()
	}
	minTime := res[0].CreatedAt
	for i := 0; i < len(res); i++ {
		f.list.List[i].ID = res[i].Id
		f.list.List[i].Title = res[i].Title
		playUrl, _ := model.NewMinioDao().GetSignedURL(res[i].PlayUrl)
		f.list.List[i].PlayURL = playUrl
		//f.list.List[i].CoverURL = res[i].CoverUrl
		f.list.List[i].CoverURL = "http://10.168.1.166:9000/videos/oriental.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=CM831I4VFIC4J0DV5X7P%2F20230912%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20230912T185431Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiJDTTgzMUk0VkZJQzRKMERWNVg3UCIsImV4cCI6MTY5NDU2NzY5NiwicGFyZW50IjoibXltaW5pb2FkbWluIn0.YCCQTR9q1Qz0oIPuahclUJXa95p2e1aluPgybQicu2ETgGPQ-bNCNWne7PAYiACH2W-0TBXKePD4RNnQXJQs1Q&X-Amz-SignedHeaders=host&versionId=null&X-Amz-Signature=142bef7b93739f4bd50f828e29f1133a8a9a93cc767db87e61e2bfb2e5cd9ced"
		f.list.List[i].FavoriteCount = res[i].FavoriteCount
		f.list.List[i].CommentCount = res[i].CommentCount
		f.list.List[i].Author = author
		f.list.List[i].IsFavorite = true
		if minTime.After(res[i].CreatedAt) {
			minTime = res[i].CreatedAt
		}
	}
	f.NextTime = minTime
	return nil, minTime
}

func QueryListInfo(lastTime time.Time) (*VideoList, time.Time, error) {
	return NewQueryListInfoFlow(lastTime).Do()
}

func NewQueryListInfoFlow(lastTime time.Time) *QueryListInfoFlow {
	return &QueryListInfoFlow{LastTime: lastTime}
}

func (f *QueryListInfoFlow) Do() (*VideoList, time.Time, error) {
	if err := f.checkParam(); err != nil {
		return nil, time.Time{}, err
	}
	if err, nextTime := f.prepareListInfo(); err != nil {
		return nil, nextTime, err
	}
	return f.list, f.NextTime, nil
}
