package model

import "gorm.io/gorm"

const TableNameSysRole = "sys_role"

// SysRole mapped from table <sys_role>
type SysRole struct {
	gorm.Model
	RoleName string     `gorm:"column:role_name;not null" json:"role_name"`
	RoleKey  string     `gorm:"column:role_key" json:"role_key"`
	RoleSort int32      `gorm:"column:role_sort;not null" json:"role_sort"`
	Type     string     `gorm:"column:type;default:1" json:"type"`
	Status   string     `gorm:"column:status;default:1" json:"status"`
	Remark   string     `gorm:"column:remark" json:"remark"`
	Depts    []*SysDept `gorm:"many2many:role_dept;" json:"depts"`
	Menus    []*SysMenu `gorm:"many2many:role_menu;" json:"menus"`
	Users    []*SysUser `gorm:"many2many:role_user;" json:"users"`
}

// TableName SysRole's table name
func (*SysRole) TableName() string {
	return TableNameSysRole
}
