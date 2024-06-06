package service

import (
	"worframe/share/core"
	"worframe/share/model"
)

type RoleService struct{}

func (s *RoleService) FindAll(page, pageSize int) ([]*model.SysRole, error) {
	var res []*model.SysRole
	err := core.DB.Limit(pageSize).Offset(page).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *RoleService) FindById(id int) (*model.SysRole, error) {
	var res *model.SysRole
	err := core.DB.Where("dept_id=?", id).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *RoleService) Create(role *model.SysRole) error {
	return core.DB.Create(&role).Error
}
func (s *RoleService) Update(role *model.SysRole) error {
	return core.DB.Save(&role).Error
}
func (s *RoleService) Delete(id int64) error {
	return core.DB.Where("role_id=?", id).Delete(&model.SysRole{}).Error
}
