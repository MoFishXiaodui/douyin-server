package model

import "testing"

func TestCreate(t *testing.T) {
	err := MySQLInit()
	if err != nil {
		t.Errorf("sth wrong happened when init MySQL: %v", err)
	}
	//通过 db 对象执行数据库操作，  db 数据库连接对象
	//并将操作的结果赋值给 user 变量 （单个用户)
	user := User{Name: "刀哥", Password: "77777"}
	_, err = NewUserDaoInstance().Create(user)
	if err != nil {
		t.Errorf("something wrong in the Create, err: %v", err)
	}

	// 重复创建
	user2 := User{Name: "刀哥", Password: "77777"}
	_, err = NewUserDaoInstance().Create(user2)
	if err == nil {
		t.Errorf("something wrong in the Create, err:%v", err)
	}

	user3 := User{Name: "小刀", Password: "99872qq!"}
	_, err = NewUserDaoInstance().Create(user3)
	if err != nil {
		t.Errorf("something wrong in the Create, err:%v", err)
	}

	user4 := User{Name: "baobao"}
	_, err = NewUserDaoInstance().Create(user4)
	if err == nil {
		t.Errorf("something wrong in the Create, err:%v", err)
	}

	user5 := User{Name: "bao", Password: "asdf46dsa4f6sa1g3sda1g32as1d1xc31@@v131gdf8564wer6q3w46qa"}
	_, err = NewUserDaoInstance().Create(user5)
	if err == nil {
		t.Errorf("something wrong in the Create, err:%v", err)
	}
}
