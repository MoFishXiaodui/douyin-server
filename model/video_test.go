package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*func TestMain(m *testing.M) {
	fmt.Println("STARTING TESTING")
	_ = MySQLInit()
	fmt.Println("Book initialized")

	os.Exit(m.Run())
}*/

func TestVideoDao_InsertNewVideo(t *testing.T) {
	_ = MySQLInit()

	var expected error = nil
	res := NewVideoDao().InsertNewVideo(123, 456, 100, 200,
		"www.baidu123.com", "www.baidu345.com", "baidu123")
	fmt.Println("res", res)
	assert.Equal(t, expected, res)
}

func TestVideoDao_InsertNewVideo2(t *testing.T) {
	var expected error = nil
	res := NewVideoDao().InsertNewVideo(345, 678, 300, 400,
		"www.baidu345.com", "www.baidu678.com", "baidu456")
	fmt.Println("res", res)
	assert.Equal(t, expected, res)
}

func TestVideoDao_QueryVideo(t *testing.T) {
	expectedtitle := "baidu123"
	video, err := NewVideoDao().QueryVideo(123)
	if err != nil {
		fmt.Println("the video", expectedtitle, "not found")
		return
	}
	assert.Equal(t, expectedtitle, video.Title)
}

func TestVideoDao_QueryVideos(t *testing.T) {
	_ = MySQLInit()

	results, _ := NewVideoDao().QueryVideos()
	fmt.Println(results)
}

func TestVideoDao_UpdateVideo(t *testing.T) {
	err := NewVideoDao().UpdateVideo(123, 456, 100, 200,
		"www.wy123.com", "www.wy456.com", "wy123")
	fmt.Println("err:", err)

}

func TestVideoDao_DeleteVideo(t *testing.T) {
	err := NewVideoDao().DeleteVideo(345)
	assert.Equal(t, nil, err)
}

func TestVideoDao_DeleteDeletedVideo(t *testing.T) {
	err := NewVideoDao().DeleteDeletedVideo(345)
	assert.Equal(t, nil, err)
}

func TestVideoDao_QueryVideosByAuthorId(t *testing.T) {
	_ = MySQLInit()
	videos, err := NewVideoDao().QueryVideosByAuthorId(3)
	if err != nil {
		fmt.Println("d")
		t.Error("unexpeted")
	}
	for _, video := range videos {
		fmt.Printf("%#v\n", video)
	}
}
