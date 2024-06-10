package migrate

import (
	"gorm.io/gorm"
	"worframe/pkg/auth/utils"
	"worframe/share/model"
)

func (m *DBMigrate) initTestData() error {

	err := testDeptInit(m.db)
	if err != nil {
		return err
	}
	err = testMenuInit(m.db)
	if err != nil {
		return err
	}
	err = testUserInit(m.db)
	if err != nil {
		return err
	}
	err = testRoleInit(m.db)
	if err != nil {
		return err
	}

	return nil
}
func testDeptInit(db *gorm.DB) error {
	depts := []*model.SysDept{
		{
			Model:    gorm.Model{ID: 1},
			DeptName: "Worframe",
			OrderNum: 0,
			Leader:   "海石花",
			Phone:    "13700000000",
			Email:    "test@test.com",
		}, {
			Model:    gorm.Model{ID: 2},
			DeptName: "管理部",
			OrderNum: 1,
			Leader:   "海石花",
			Phone:    "13700000000",
			Email:    "test@test.com",
		}, {
			Model:    gorm.Model{ID: 3},
			DeptName: "项目部",
			OrderNum: 2,
			Leader:   "海石花",
			Phone:    "13700000000",
			Email:    "test@test.com",
		}, {
			Model:    gorm.Model{ID: 4},
			DeptName: "开发部",
			OrderNum: 3,
			Leader:   "海石花",
			Phone:    "13700000000",
			Email:    "test@test.com",
		}, {
			Model:    gorm.Model{ID: 5},
			DeptName: "市场部",
			OrderNum: 4,
			Leader:   "海石花",
			Phone:    "13700000000",
			Email:    "test@test.com",
		}, {
			Model:    gorm.Model{ID: 6},
			DeptName: "测试部",
			OrderNum: 5,
			Leader:   "海石花",
			Phone:    "13700000000",
			Email:    "test@test.com",
		}, {
			Model:    gorm.Model{ID: 7},
			DeptName: "运营部",
			OrderNum: 6,
			Leader:   "海石花",
			Phone:    "13700000000",
			Email:    "test@test.com",
		}, {
			Model:    gorm.Model{ID: 8},
			DeptName: "财务部",
			OrderNum: 7,
			Leader:   "海石花",
			Phone:    "13700000000",
			Email:    "test@test.com",
		}, {
			Model:    gorm.Model{ID: 9},
			DeptName: "人事部",
			OrderNum: 1,
			Leader:   "海石花",
			Phone:    "13700000000",
			Email:    "test@test.com",
		},
	}

	return db.Save(&depts).Error
}
func testRoleInit(db *gorm.DB) error {
	roles := []model.SysRole{
		{
			Model:    gorm.Model{ID: 1},
			RoleName: "role_public",
			RoleSort: 1,
			Type:     "1",
		},
		{
			Model:    gorm.Model{ID: 2},
			RoleName: "role_default",
			RoleSort: 2,
			Type:     "1",
		},
		{
			Model:    gorm.Model{ID: 3},
			RoleName: "role_admin",
			RoleSort: 3,
			Type:     "1",
			Menus: []*model.SysMenu{
				{Model: gorm.Model{ID: 1}},
			},
			Depts: []*model.SysDept{
				{Model: gorm.Model{ID: 4}},
			},
		},
		{
			Model:    gorm.Model{ID: 4},
			RoleName: "role_dept_all",
			RoleSort: 4,
			Type:     "1", Menus: []*model.SysMenu{
				{Model: gorm.Model{ID: 2}},
				{Model: gorm.Model{ID: 3}},
				{Model: gorm.Model{ID: 4}},
				{Model: gorm.Model{ID: 5}},
				{Model: gorm.Model{ID: 6}},
			},
		},
		{
			Model:    gorm.Model{ID: 5},
			RoleName: "role_role_all",
			RoleSort: 5,
			Type:     "1", Menus: []*model.SysMenu{
				{Model: gorm.Model{ID: 7}},
				{Model: gorm.Model{ID: 8}},
				{Model: gorm.Model{ID: 9}},
				{Model: gorm.Model{ID: 10}},
				{Model: gorm.Model{ID: 11}},
			},
		},
		{
			Model:    gorm.Model{ID: 6},
			RoleName: "role_menu_all",
			RoleSort: 6,
			Type:     "1", Menus: []*model.SysMenu{
				{Model: gorm.Model{ID: 12}},
				{Model: gorm.Model{ID: 13}},
				{Model: gorm.Model{ID: 14}},
				{Model: gorm.Model{ID: 15}},
				{Model: gorm.Model{ID: 16}},
			},
		},
		{
			Model:    gorm.Model{ID: 7},
			RoleName: "人事主管",
			RoleSort: 7,
			Type:     "1", Menus: []*model.SysMenu{
				{Model: gorm.Model{ID: 4}},
				{Model: gorm.Model{ID: 5}},
				{Model: gorm.Model{ID: 6}},
			},
		},
		{
			Model:    gorm.Model{ID: 8},
			RoleName: "人事专员",

			RoleSort: 7,
			Type:     "1", Menus: []*model.SysMenu{
				{Model: gorm.Model{ID: 2}},
				{Model: gorm.Model{ID: 3}},
				{Model: gorm.Model{ID: 7}},
				{Model: gorm.Model{ID: 8}},
				{Model: gorm.Model{ID: 12}},
				{Model: gorm.Model{ID: 13}},
			},
			Depts: []*model.SysDept{
				{Model: gorm.Model{ID: 9}},
			},
		},
	}
	return db.Save(roles).Error
}

func testMenuInit(db *gorm.DB) error {
	menuData := []model.SysMenu{
		{Model: gorm.Model{ID: 1}, MenuName: "超级管理员",
			ParentID: 0, OrderNum: 1, URL: "/*", Method: "*", MenuType: "M", Visible: "0", IsRefresh: "1", Perms: "system:admin", Icon: "fa fa-gear", Remark: "admin"},
		{Model: gorm.Model{ID: 2}, MenuName: "dept:getAll",
			ParentID: 0, OrderNum: 1, URL: "/dept", Method: "GET", MenuType: "M", Visible: "0", IsRefresh: "1", Perms: "dept:getAll", Icon: "fa fa-gear", Remark: "admin"},
		{Model: gorm.Model{ID: 3}, MenuName: "dept:getOne",
			ParentID: 0, OrderNum: 2, URL: "/dept/:id", Method: "GET", MenuType: "M", Visible: "0", IsRefresh: "1", Perms: "dept:getOne", Icon: "fa fa-video-camera", Remark: "admin"},
		{Model: gorm.Model{ID: 4}, MenuName: "dept:post",
			ParentID: 0, OrderNum: 3, URL: "/dept", Method: "POST", MenuType: "M", Visible: "0", IsRefresh: "1", Perms: "dept:post", Icon: "fa fa-bars", Remark: "admin"},
		{Model: gorm.Model{ID: 5}, MenuName: "dept:put",
			ParentID: 0, OrderNum: 4, URL: "/dept/:id", Method: "PUT", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "dept:put", Icon: "fa fa-location-arrow", Remark: "admin"},
		{Model: gorm.Model{ID: 6}, MenuName: "dept:delete",
			ParentID: 0, OrderNum: 1, URL: "/dept/:id", Method: "DELETE", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "dept:delete", Icon: "fa fa-user-o", Remark: "admin"},
		{Model: gorm.Model{ID: 7}, MenuName: "role:getAll",
			ParentID: 0, OrderNum: 2, URL: "/role", Method: "GET", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "role:getAll", Icon: "fa fa-user-secret", Remark: "admin"},
		{Model: gorm.Model{ID: 8}, MenuName: "role:getOne",
			ParentID: 0, OrderNum: 3, URL: "/role/:id", Method: "GET", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "role:getOne", Icon: "fa fa-th-list", Remark: "admin"},
		{Model: gorm.Model{ID: 9}, MenuName: "role:post",
			ParentID: 0, OrderNum: 4, URL: "/role", Method: "POST", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "role:post", Icon: "fa fa-outdent", Remark: "admin"},
		{Model: gorm.Model{ID: 10}, MenuName: "role:put",
			ParentID: 0, OrderNum: 5, URL: "/role/:id", Method: "PUT", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "role:put", Icon: "fa fa-address-card-o", Remark: "admin"},
		{Model: gorm.Model{ID: 11}, MenuName: "role:delete",
			ParentID: 0, OrderNum: 6, URL: "/role/:id", Method: "DELETE", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "role:delete", Icon: "fa fa-bookmark-o", Remark: "admin"},
		{Model: gorm.Model{ID: 12}, MenuName: "menu:getAll",
			ParentID: 0, OrderNum: 7, URL: "/menu", Method: "GET", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "menu:getAll", Icon: "fa fa-sun-o", Remark: "admin"},
		{Model: gorm.Model{ID: 13}, MenuName: "menu:getOne",
			ParentID: 0, OrderNum: 8, URL: "/menu/:id", Method: "GET", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "menu:getOne", Icon: "fa fa-bullhorn", Remark: "admin"},
		{Model: gorm.Model{ID: 14}, MenuName: "menu:post",
			ParentID: 0, OrderNum: 9, URL: "/menu", Method: "POST", MenuType: "M", Visible: "0", IsRefresh: "1", Perms: "menu:post", Icon: "fa fa-pencil-square-o", Remark: "admin"},
		{Model: gorm.Model{ID: 15}, MenuName: "menu:put",
			ParentID: 0, OrderNum: 1, URL: "/menu/:id", Method: "PUT", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "menu:put", Icon: "fa fa-user-circle", Remark: "admin"},
		{Model: gorm.Model{ID: 16}, MenuName: "menu:delete",
			ParentID: 0, OrderNum: 2, URL: "/menu/:id", Method: "DELETE", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "menu:delete", Icon: "fa fa-tasks", Remark: "admin"},
	}
	return db.Save(menuData).Error
}
func testUserInit(db *gorm.DB) error {
	user := []model.SysUser{
		{Model: gorm.Model{ID: 1}, DeptId: 2, UserType: "00", UserName: "admin", NickName: "超级管理员"},
		{Model: gorm.Model{ID: 2}, DeptId: 1, UserType: "00", UserName: "user_one", NickName: "普通用户"},
	}
	pwd := "123456"
	user[0].Salt = utils.SaltSpawn()
	user[1].Salt = utils.SaltSpawn()
	user[0].Password = utils.EncryptPassword(user[0].UserName, pwd, user[0].Salt)
	user[1].Password = utils.EncryptPassword(user[1].UserName, pwd, user[1].Salt)
	return db.Save(user).Error
}
