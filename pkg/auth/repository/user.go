package repository

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
	"worframe/pkg/auth/core/iface"
	"worframe/pkg/auth/event"
	"worframe/share/model"
)

type UserRepo struct {
	core  *iface.ICore
	db    *gorm.DB
	redis *redis.Pool
}

func NewUserRepository(core iface.ICore) *UserRepo {
	return &UserRepo{
		core:  &core,
		db:    core.GetDB(),
		redis: core.GetRedis(),
	}
}
func (r *UserRepo) FindUserByUsername(username string) (*model.SysUser, error) {
	var User model.SysUser
	err := r.db.Model(&model.SysUser{}).Where(&model.SysUser{UserName: username}).First(&User).Error
	return &User, err
}
func (r *UserRepo) RegisterUser(user model.SysUser) (*model.SysUser, error) {
	/**
	检测用户名是否合法,以及用户名类型
	*/
	usernameType, ok := event.ValidateUsername(user.UserName)
	if !ok {
		return nil, fmt.Errorf("invalid username: %s", user.UserName)
	}
	switch usernameType {
	case "phone":
		user.Phone = user.UserName
	case "email":
		user.Email = user.UserName
	}
	/**
	检测用户是否存在
	*/
	var ExistUser model.SysUser
	r.db.Where("phone = ? OR email = ? OR user_name = ?", user.UserName, user.UserName, user.UserName).First(&ExistUser)
	if ExistUser.ID != 0 {
		return nil, fmt.Errorf("user already exists")
	}
	/**
	如果存在密码则加密密码
	*/
	if user.Password != "" {
		password, salt, err := event.CreateNewPassword(user.Password)
		if err != nil {
			return nil, err
		}
		user.Password = password
		user.Salt = salt
	}
	/**
	保存
	*/
	return &user, r.db.Create(&user).Error
}
