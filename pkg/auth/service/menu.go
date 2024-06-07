package service

import (
	"worframe/share/core"
	"worframe/share/model"
)

type MenuService struct{}

func (s *MenuService) FindAll(page, pageSize int) ([]*model.SysMenu, error) {
	var res []*model.SysMenu
	err := core.DB.Limit(pageSize).Offset((page - 1) * pageSize).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *MenuService) FindById(id uint) (*model.SysMenu, error) {
	var res *model.SysMenu
	err := core.DB.Where(id).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *MenuService) Create(menu *model.SysMenu) error {
	return core.DB.Create(&menu).Error
}
func (s *MenuService) Update(menu *model.SysMenu) error {
	return core.DB.Updates(&menu).Error
}
func (s *MenuService) Delete(id uint) error {
	return core.DB.Where(id).Delete(&model.SysMenu{}).Error
}
