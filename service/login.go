package service

import (
	"dy/model"
	"errors"
	"fmt"
)

type UserInfo struct {
	Token  string
	State  bool
	UserId uint
}
type UserLoginFlow struct {
	UserName string
	Password string
	UserInfo *UserInfo
}

// prepareInfo用来通过底层模型获取对应数据(token)的函数
func (f *UserLoginFlow) prepareInfo() error {
	user := model.NewUserDaoInstance().QuerywithName(f.UserName)
	if user == nil {
		return errors.New("This user does not exist")
	}
	f.UserInfo.Token = user.Token
	f.UserInfo.UserId = user.ID
	return nil
}

// NewUserQueryLoginFlow 创建新UserQueryLoginFlow实例的函数
func NewUserLoginFlow(name, password string) *UserLoginFlow {
	return &UserLoginFlow{UserName: name, Password: password}
}

func (f *UserLoginFlow) DO() *UserInfo {
	if len(f.Password) > 32 {
		f.UserInfo.State = false
		return f.UserInfo
	}
	f.UserInfo = &UserInfo{}
	user := model.NewUserDaoInstance().QuerywithNameAndPassword(f.UserName, f.Password)
	fmt.Println(f.UserName, f.UserInfo.State, "tt")

	if user == nil {
		f.UserInfo.State = false
		return f.UserInfo
	}

	if err := f.prepareInfo(); err != nil {
		f.UserInfo.State = false
		return f.UserInfo
	}
	f.UserInfo.State = true
	return f.UserInfo
}

// UserLogin 外部来了一个请求，我们创建一个新的UserLoginFlow结构来处理，这个结构的Do()方法最终返回想要的结果
func UserLogin(name, password string) *UserInfo {

	return NewUserLoginFlow(name, password).DO()
}
