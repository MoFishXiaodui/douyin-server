package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFavoriteDao_InsertNewFavorite(t *testing.T) {
	var expected error = nil
	res := NewFavoriteDao().InsertNewFavorite(123, "wjl", true)
	fmt.Println("res", res)
	assert.Equal(t, expected, res)
}

func TestFavoriteDao_QueryFavorite(t *testing.T) {
	expectedlover := "wjl"
	favorite, err := NewFavoriteDao().QueryFavorite(123)
	if err != nil {
		fmt.Println("the favorite", expectedlover, "not found")
		return
	}
	assert.Equal(t, expectedlover, favorite.Lover)
}

func TestFavoriteDao_UpdateFavorite(t *testing.T) {
	err := NewFavoriteDao().UpdateFavorite(123, "lzn", true)
	fmt.Println("err", err)
}

func TestFavoriteDao_DeleteFavorite(t *testing.T) {
	err := NewFavoriteDao().DeleteFavorite(123)
	assert.Equal(t, nil, err)
}

func TestFavoriteDao_DeleteDeleteFavorite(t *testing.T) {
	err := NewFavoriteDao().DeleteDeleteFavorite(123)
	assert.Equal(t, nil, err)
}
