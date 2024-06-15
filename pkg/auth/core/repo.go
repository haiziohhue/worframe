package core

import (
	"errors"
	"worframe/pkg/auth/core/iface"
	"worframe/pkg/auth/repository"
)

func (ac *AuthCore) InitRepository() iface.ICore {
	ac.Repo["dept"] = repository.NewDeptRepository(ac)
	ac.Repo["user"] = repository.NewUserRepository(ac)
	return ac
}
func (ac *AuthCore) GetRepo(name string) (interface{}, error) {
	repo := ac.Repo[name]
	if repo == nil {
		return nil, errors.New("repo Not Found")
	}
	return repo, nil
}
