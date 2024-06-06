package service

import (
	"worframe/share/core"
	"worframe/share/model"
)

type DeptService struct{}

func (s *DeptService) FindAll(page, pageSize int) ([]*model.SysDept, error) {
	var res []*model.SysDept
	err := core.DB.Limit(pageSize).Offset(page).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *DeptService) FindById(id int) (*model.SysDept, error) {
	var res *model.SysDept
	err := core.DB.Where(id).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *DeptService) Create(dept *model.SysDept) error {
	return core.DB.Create(&dept).Error
}
func (s *DeptService) Update(dept *model.SysDept) error {
	return core.DB.Save(&dept).Error
}
func (s *DeptService) Delete(id int64) error {
	return core.DB.Where(id).Delete(&model.SysDept{}).Error
}
