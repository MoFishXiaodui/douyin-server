package service

import (
	"dy/model"
	"errors"
)

type UserInfo struct {
	Token string
	State bool
}
type UserQueryLoginFlow struct {
	Name     string
	Password string
	UserInfo *UserInfo
}

// prepareInfo用来通过底层模型获取对应数据(token)的函数
func (f *UserQueryLoginFlow) prepareInfo() error {
	user := model.NewUserDaoInstance().QuerywithName(f.Name)
	if user == nil {
		return errors.New("This user does not exist")
	}
	f.UserInfo.Token = user.Token
	return nil
}

// NewUserQueryLoginFlow 创建新UserQueryLoginFlow实例的函数
func NewUserQueryLoginFlow(name, password string) *UserQueryLoginFlow {
	return &UserQueryLoginFlow{Name: name, Password: password}
}

func (f *UserQueryLoginFlow) DO() *UserInfo {
	if len(f.Password) > 32 {
		f.UserInfo.State = false
		return f.UserInfo
	}
	f.UserInfo = &UserInfo{}

	user := model.NewUserDaoInstance().QuerywithNameAndPassword(f.Name, f.Password)
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

// UserQueryLogin 外部来了一个请求，我们创建一个新的UserQueryLoginFlow结构来处理，这个结构的Do()方法最终返回想要的结果
func UserQueryLogin(name, password string) *UserInfo {

	return NewUserQueryLoginFlow(name, password).DO()
}
