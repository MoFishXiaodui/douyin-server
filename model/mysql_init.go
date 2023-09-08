package model

import (
	"dy/config"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func MySQLInit() error {
	user, pwd, addr, dbName := config.GetMySQLConfig()
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, addr, dbName)
	dbTemp, err := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic("MySQL数据库连接失败")
	}
	db = dbTemp

	err = dbMigrate()
	if err != nil {
		panic("数据库初始化表格失败: " + err.Error())
	}

	return nil
}

func GetMySQLdb() *gorm.DB {
	dbOnce.Do(func() {

	})
	return db
}

func dbMigrate() error {
	err1 := UserInit()
	err2 := RelationInit()
	if err1 != nil {
		return errors.New("初始化用户表失败:" + err1.Error())
	}
	if err2 != nil {
		return errors.New("初始化用户表失败:" + err2.Error())
	}

	err := InitFavorite()
	if err1 != nil {
		return errors.New("初始化视频喜好失败" + err.Error())
	}

	err = InitVideo()
	if err2 != nil {
		return errors.New("初始化视频信息失败" + err.Error())
	}
	return nil
}
