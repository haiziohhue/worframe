package migrate

import (
	"worframe/share/model"
)

type roleDept struct {
}

func (*roleDept) TableName() string {
	return "role_dept"
}

type roleMenu struct{}

func (*roleMenu) TableName() string {
	return "role_menu"
}

type roleUser struct {
}

func (*roleUser) TableName() string {
	return "role_user"
}
func (m *DBMigrate) initTable() error {
	err := m.db.Migrator().DropTable(
		&roleDept{},
		&roleMenu{},
		&roleUser{},
		&model.SysUser{},
		&model.SysDept{},
		&model.SysMenu{},
		&model.SysRole{},
	)
	if err != nil {
		return err
	}
	err = m.db.AutoMigrate(
		&model.SysUser{},
		&model.SysDept{},
		&model.SysMenu{},
		&model.SysRole{},
	)
	if err != nil {
		return err
	}
	return nil
}
