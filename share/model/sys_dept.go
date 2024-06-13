package model

import (
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

const TableNameSysDept = "sys_dept"

type SysDept struct {
	gorm.Model
	ParentID  int64      `gorm:"column:parent_id" json:"parent_id" validate:""`
	Ancestors string     `gorm:"column:ancestors" json:"ancestors" validate:""`
	DeptName  string     `gorm:"column:dept_name" json:"dept_name" validate:""`
	OrderNum  int32      `gorm:"column:order_num" json:"order_num"validate:""`
	Leader    string     `gorm:"column:leader" json:"leader" validate:"require"`
	Phone     string     `gorm:"column:phone" json:"phone" validate:""`
	Email     string     `gorm:"column:email" json:"email" validate:""`
	Status    string     `gorm:"column:status;default:0" json:"status" validate:""`
	Roles     []*SysRole `gorm:"many2many:role_dept" json:"roles" validate:""`
	User      []SysUser  `gorm:"foreignKey:DeptId;" json:"user" validate:""`
}

// TableName SysDept's table name
func (*SysDept) TableName() string {
	return TableNameSysDept
}
func (m *SysDept) Validate(model SysDept) error {
	validate := validator.New()
	return validate.Struct(model)
}
