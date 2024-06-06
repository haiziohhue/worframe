package model

import "gorm.io/gorm"

const TableNameSysMenu = "sys_menu"

// SysMenu mapped from table <sys_menu>
type SysMenu struct {
	gorm.Model
	MenuName  string    `gorm:"column:menu_name;not null" json:"menu_name"`
	ParentID  int64     `gorm:"column:parent_id" json:"parent_id"`
	OrderNum  int32     `gorm:"column:order_num" json:"order_num"`
	URL       string    `gorm:"column:url;default:#" json:"url"`
	Target    string    `gorm:"column:target" json:"target"`
	MenuType  string    `gorm:"column:menu_type" json:"menu_type"`
	Visible   string    `gorm:"column:visible;default:0" json:"visible"`
	IsRefresh string    `gorm:"column:is_refresh;default:1" json:"is_refresh"`
	Perms     string    `gorm:"column:perms" json:"perms"`
	Icon      string    `gorm:"column:icon;default:#" json:"icon"`
	Remark    string    `gorm:"column:remark" json:"remark"`
	Roles     []SysRole `gorm:"many2many:role_menu" json:"roles"`
}

// TableName SysMenu's table name
func (*SysMenu) TableName() string {
	return TableNameSysMenu
}
