package model

import (
	"dy/config"
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

	return nil
}

func GetMySQLdb() *gorm.DB {
	dbOnce.Do(func() {

	})
	return db
}
