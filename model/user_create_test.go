package model

import "testing"

func TestCreate(t *testing.T) {
	err := MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
	//通过 db 对象执行数据库操作，  db 数据库连接对象
	//并将操作的结果赋值给 user 变量 （单个用户)
	user := User{Name: "刀哥"}
	_, err = NewUserDaoIstance().Create(user)
	if err != nil {
		t.Error("something wrong in the Create, err: ")
	}
}
