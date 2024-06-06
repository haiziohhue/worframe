package model

import "gorm.io/gorm"

const TableNameSysDept = "sys_dept"

type SysDept struct {
	gorm.Model
	ParentID  int64     `gorm:"column:parent_id" json:"parent_id"`
	Ancestors string    `gorm:"column:ancestors" json:"ancestors"`
	DeptName  string    `gorm:"column:dept_name" json:"dept_name"`
	OrderNum  int32     `gorm:"column:order_num" json:"order_num"`
	Leader    string    `gorm:"column:leader" json:"leader"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Email     string    `gorm:"column:email" json:"email"`
	Status    string    `gorm:"column:status;default:0" json:"status"`
	Roles     []SysRole `gorm:"many2many:role_dept" json:"roles"`
	User      []SysUser `gorm:"foreignKey:DeptId;" json:"user"`
}

// TableName SysDept's table name
func (*SysDept) TableName() string {
	return TableNameSysDept
}
