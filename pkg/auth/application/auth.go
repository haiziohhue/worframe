package application

import (
	"fmt"
	"worframe/pkg/auth/core/iface"
	"worframe/pkg/auth/event"
	"worframe/pkg/auth/repository"
	"worframe/share/model"
	"worframe/share/types"
)

type LoginParams struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Code     string `form:"code"`
	CodeId   string `form:"code_id"`
	Email    string `form:"email"`
	Mobile   string `form:"mobile"`
}
type AuthService struct {
	Core     *iface.ICore
	UserRepo repository.UserRepo
}

func NewAuthService(core *iface.ICore) (*AuthService, error) {
	getRepo, err := (*core).GetRepo("user")
	if err != nil {
		return nil, err
	}
	repo, ok := getRepo.(*repository.UserRepo)
	if !ok {
		return nil, fmt.Errorf("NewAuthService: Repo cast failed")
	}

	return &AuthService{core, *repo}, nil
}

func (s *AuthService) passwordLogin(params LoginParams) (*model.SysUser, error) {
	if params.Username == "" || params.Password == "" {
		return nil, fmt.Errorf("invalid params")
	}
	user, err := s.UserRepo.FindUserByUsername(params.Username)
	if err != nil {
		return nil, err
	}
	if ok := event.ComparePassword(params.Password, user); !ok {
		return nil, fmt.Errorf("invalid password")
	}
	return user, nil
}
func (s *AuthService) Login(params LoginParams, method string) (*types.JwtPayload, error) {

	var user *model.SysUser

	switch method {
	case "password":
		u, err := s.passwordLogin(params)
		if err != nil {
			return nil, err
		}
		user = u
	}

	if user == nil {
		return nil, fmt.Errorf("invalid params")
	}
	roles := event.ExportUserRole(*user)
	return &types.JwtPayload{
		UUID: user.UUID,
		Role: roles,
	}, nil
}
func (s *AuthService) Register(params LoginParams, method string) error {
	if method == "password" && params.Password == "" {
		return fmt.Errorf("invalid params")
	}
	user := &model.SysUser{
		UserName: params.Username,
		Password: params.Password,
		Email:    params.Email,
		Phone:    params.Mobile,
	}
	_, err := s.UserRepo.RegisterUser(*user)
	if err != nil {
		return err
	}
	return nil
}
