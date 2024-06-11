package service

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"worframe/share/model"
)

type DeptService struct {
	Logger *zap.Logger
	DB     *gorm.DB
}

func NewDeptService(zap *zap.Logger, db *gorm.DB) *DeptService {
	return &DeptService{
		Logger: zap,
		DB:     db,
	}
}
func (s *DeptService) FindAll(page, pageSize int) ([]*model.SysDept, error) {
	var res []*model.SysDept
	err := s.DB.Limit(pageSize).Offset((page - 1) * pageSize).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *DeptService) FindById(id uint) (*model.SysDept, error) {
	var res *model.SysDept
	err := s.DB.Where(id).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *DeptService) Create(dept *model.SysDept) error {
	return s.DB.Create(&dept).Error
}
func (s *DeptService) Update(dept *model.SysDept) error {
	return s.DB.Updates(&dept).Error
}
func (s *DeptService) Delete(id uint) error {
	return s.DB.Where(id).Delete(&model.SysDept{}).Error
}
