package strategy

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"worframe/share/core"
	"worframe/share/model"
)

func rbacWithPattern(db *gorm.DB) error {
	var menus []model.SysMenu
	var roles []model.SysRole

	if err := db.Find(&menus).Error; err != nil {
		core.Logger.Error(err)
	}
	if err := db.Find(&roles).Error; err != nil {
		core.Logger.Error(err)
	}

	// menus record
	rules := make([]gormadapter.CasbinRule, 0)
	for _, menu := range menus {
		rule := &gormadapter.CasbinRule{
			Ptype: "p",
			V0:    menu.Perms,
			V1:    menu.URL,
			V2:    menu.Method,
		}
		rules = append(rules, *rule)
	}
	// role record
	for _, role := range roles {
		if role.Menus != nil {
			for _, menu := range role.Menus {
				rule := &gormadapter.CasbinRule{
					Ptype: "g",
					V0:    role.RoleName,
					V1:    menu.Perms,
				}
				rules = append(rules, *rule)
			}
		}
		if role.Depts != nil {
			for _, dept := range role.Depts {
				rule := &gormadapter.CasbinRule{
					Ptype: "g",
					V0:    dept.DeptName,
					V1:    role.RoleName,
				}
				rules = append(rules, *rule)
			}
		}
		if role.Users != nil {
			for _, user := range role.Users {
				rule := &gormadapter.CasbinRule{
					Ptype: "g",
					V0:    user.UserName,
					V1:    role.RoleName,
				}
				rules = append(rules, *rule)
			}
		}
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		tx.Migrator().DropTable(&gormadapter.CasbinRule{}).Error()
		tx.Create(&rules)
		return nil
	})
	if err != nil {
		core.Logger.Error(err)
	}
	return nil
}