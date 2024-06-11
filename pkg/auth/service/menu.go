package service

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"worframe/share/model"
)

type MenuService struct {
	Logger *zap.Logger
	DB     *gorm.DB
}

func NewMenuService(zap *zap.Logger, db *gorm.DB) *MenuService {
	return &MenuService{Logger: zap,
		DB: db}
}
func (s *MenuService) FindAll(page, pageSize int) ([]*model.SysMenu, error) {
	var res []*model.SysMenu
	err := s.DB.Limit(pageSize).Offset((page - 1) * pageSize).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *MenuService) FindById(id uint) (*model.SysMenu, error) {
	var res *model.SysMenu
	err := s.DB.Where(id).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *MenuService) Create(menu *model.SysMenu) error {
	return s.DB.Create(&menu).Error
}
func (s *MenuService) Update(menu *model.SysMenu) error {
	return s.DB.Updates(&menu).Error
}
func (s *MenuService) Delete(id uint) error {
	return s.DB.Where(id).Delete(&model.SysMenu{}).Error
}
