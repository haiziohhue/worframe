package repository

import (
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
	"worframe/pkg/auth/core/iface"
	"worframe/share/model"
)

type DeptRepository struct {
	core  *iface.ICore
	db    *gorm.DB
	redis *redis.Pool
}

func NewDeptRepository(core iface.ICore) *DeptRepository {
	return &DeptRepository{
		&core, core.GetDB(), core.GetRedis(),
	}
}
func (d *DeptRepository) Create(dept model.SysDept) error {
	return d.db.Create(&dept).Error
}
func (d *DeptRepository) Update(dept model.SysDept) error {
	return d.db.Updates(&dept).Error
}
func (d *DeptRepository) Delete(dept model.SysDept) error {
	return d.db.Delete(&dept).Error
}
func (d *DeptRepository) Find(page, pageSize int) ([]model.SysDept, error) {
	var depts []model.SysDept
	err := d.db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&depts).Error
	if err != nil {
		return nil, err
	}
	return depts, nil
}
func (d *DeptRepository) FindOne(dept model.SysDept) (model.SysDept, error) {
	res := model.SysDept{}
	err := d.db.First(&dept).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
func (d *DeptRepository) FindWhere(dept model.SysDept) ([]model.SysDept, error) {
	var depts []model.SysDept
	if err := d.db.Where(&dept).Find(&depts).Error; err != nil {
		return nil, err
	}
	return depts, nil
}
