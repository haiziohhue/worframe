package model

import "gorm.io/gorm"

const TableNameSysUser = "sys_user"

// SysUser mapped from table <sys_user>
type SysUser struct {
	gorm.Model
	UUID     string `gorm:"column:uuid;default:uuid_generate_v4()" json:"uuid"`
	UserType string `gorm:"column:user_type;default:00" json:"user_type"`
	UserName string `gorm:"column:user_name" json:"user_name"`
	NickName string `gorm:"column:nick_name;not null" json:"nick_name"`
	Email    string `gorm:"column:email" json:"email"`
	Phone    string `gorm:"column:phone" json:"phone"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
	Password string `gorm:"column:password" json:"password"`
	Salt     string `gorm:"column:salt" json:"salt"`
	Status   string `gorm:"column:status;default:0" json:"status"`
	Remark   string `gorm:"column:remark" json:"remark"`
	DeptId   int64  `gorm:"column:dept_id;default:0" json:"dept_id"`
	//Dept     *SysDept
	Role []*SysRole `gorm:"many2many:role_user;" json:"roles"`
}

// TableName SysUser's table name
func (*SysUser) TableName() string {
	return TableNameSysUser
}
