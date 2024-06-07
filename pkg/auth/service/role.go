package service

import (
	"worframe/share/core"
	"worframe/share/model"
)

type RoleService struct{}

func (s *RoleService) FindAll(page, pageSize int) ([]*model.SysRole, error) {
	var res []*model.SysRole
	err := core.DB.Limit(pageSize).Offset((page - 1) * pageSize).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *RoleService) FindById(id uint) (*model.SysRole, error) {
	var res *model.SysRole
	err := core.DB.Where(id).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *RoleService) Create(role *model.SysRole) error {
	return core.DB.Create(&role).Error
}
func (s *RoleService) Update(role *model.SysRole) error {
	return core.DB.Updates(&role).Error
}
func (s *RoleService) Delete(id uint) error {
	return core.DB.Where(id).Delete(&model.SysRole{}).Error
}
