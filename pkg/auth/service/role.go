package service

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"worframe/share/model"
)

type RoleService struct {
	Logger *zap.Logger
	DB     *gorm.DB
}

func NewRoleService(zap *zap.Logger, db *gorm.DB) *RoleService {
	return &RoleService{Logger: zap,
		DB: db}
}
func (s *RoleService) FindAll(page, pageSize int) ([]*model.SysRole, error) {
	var res []*model.SysRole
	err := s.DB.Limit(pageSize).Offset((page - 1) * pageSize).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *RoleService) FindById(id uint) (*model.SysRole, error) {
	var res *model.SysRole
	err := s.DB.Where(id).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *RoleService) Create(role *model.SysRole) error {
	return s.DB.Create(&role).Error
}
func (s *RoleService) Update(role *model.SysRole) error {
	return s.DB.Updates(&role).Error
}
func (s *RoleService) Delete(id uint) error {
	return s.DB.Where(id).Delete(&model.SysRole{}).Error
}
