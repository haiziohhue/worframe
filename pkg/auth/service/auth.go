package service

import (
	"worframe/share/core"
	"worframe/share/model"
)

type AuthService struct {
}

func (s *AuthService) Login(username, password string) (error, string) {
	var user model.SysUser
	findUser := core.DB.Where("username = ?", username).First(&user)
	if findUser.Error != nil {
		return findUser.Error, ""
	}
	if user.Password != password {
		return findUser.Error, ""
	}
	return nil, ""
}
