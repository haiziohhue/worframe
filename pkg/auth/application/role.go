package application

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"worframe/pkg/auth/core/iface"
	"worframe/share/model"
)

type RoleService struct {
	Core   *iface.ICore
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewRoleService(core iface.ICore) *RoleService {
	return &RoleService{Core: &core,
		DB:     core.GetDB(),
		Logger: core.GetLog()}
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
