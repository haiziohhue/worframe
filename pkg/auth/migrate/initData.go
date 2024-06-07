package migrate

import (
	"gorm.io/gorm"
	"worframe/share/model"
	"worframe/share/utils"
)

func (m *DBMigrate) initTestData() error {
	err := deptInit(m.db)
	if err != nil {
		return err
	}
	err = roleInit(m.db)
	if err != nil {
		return err
	}
	err = userInit(m.db)
	if err != nil {
		return err
	}
	err = menuInit(m.db)
	if err != nil {
		return err
	}
	err = bindInit(m.db)
	if err != nil {
		return err
	}
	return nil
}
func deptInit(db *gorm.DB) error {
	depts := []*model.SysDept{
		{Model: gorm.Model{ID: 100},
			ParentID:  0,
			Ancestors: "0",
			DeptName:  "项目总部",
			OrderNum:  0,
			Leader:    "海石花",
			Phone:     "13700000000",
			Email:     "test@test.com",
		}, {Model: gorm.Model{ID: 101},
			ParentID:  100,
			Ancestors: "0,100",
			DeptName:  "广州分公司",
			OrderNum:  1,
			Leader:    "海石花",
			Phone:     "13700000000",
			Email:     "test@test.com",
		}, {Model: gorm.Model{ID: 102},
			ParentID:  100,
			Ancestors: "0,100",
			DeptName:  "上海分公司",
			OrderNum:  2,
			Leader:    "海石花",
			Phone:     "13700000000",
			Email:     "test@test.com",
		}, {Model: gorm.Model{ID: 103},
			ParentID:  101,
			Ancestors: "0,100,101",
			DeptName:  "研发部门",
			OrderNum:  1,
			Leader:    "海石花",
			Phone:     "13700000000",
			Email:     "test@test.com",
		}, {Model: gorm.Model{ID: 104},
			ParentID:  101,
			Ancestors: "0,100,101",
			DeptName:  "市场部门",
			OrderNum:  2,
			Leader:    "海石花",
			Phone:     "13700000000",
			Email:     "test@test.com",
		}, {Model: gorm.Model{ID: 105},
			ParentID:  101,
			Ancestors: "0,100,101",
			DeptName:  "测试部门",
			OrderNum:  3,
			Leader:    "海石花",
			Phone:     "13700000000",
			Email:     "test@test.com",
		}, {Model: gorm.Model{ID: 106},
			ParentID:  101,
			Ancestors: "0,100,101",
			DeptName:  "运维部门",
			OrderNum:  5,
			Leader:    "海石花",
			Phone:     "13700000000",
			Email:     "test@test.com",
		}, {Model: gorm.Model{ID: 107},
			ParentID:  101,
			Ancestors: "0,100,101",
			DeptName:  "财务部门",
			OrderNum:  4,
			Leader:    "海石花",
			Phone:     "13700000000",
			Email:     "test@test.com",
		}, {Model: gorm.Model{ID: 108},
			ParentID:  102,
			Ancestors: "0,100,102",
			DeptName:  "市场部门",
			OrderNum:  1,
			Leader:    "海石花",
			Phone:     "13700000000",
			Email:     "test@test.com",
		}, {Model: gorm.Model{ID: 109},
			ParentID:  102,
			Ancestors: "0,100,102",
			DeptName:  "财务部门",
			OrderNum:  2,
			Leader:    "海石花",
			Phone:     "13700000000",
			Email:     "test@test.com",
		},
	}

	return db.Save(&depts).Error
}
func roleInit(db *gorm.DB) error {
	roles := []model.SysRole{
		{
			RoleName: "admin",
			RoleSort: 1,
			Type:     "1",
		},
		{
			RoleName: "default",
			RoleSort: 2,
			Type:     "1",
		},
	}
	return db.Save(roles).Error
}
func menuInit(db *gorm.DB) error {
	menuData := []model.SysMenu{
		{Model: gorm.Model{ID: 1}, MenuName: "系统管理", ParentID: 0, OrderNum: 1, URL: "#", Target: "", MenuType: "M", Visible: "0", IsRefresh: "1", Perms: "", Icon: "fa fa-gear", Remark: "admin"},
		{Model: gorm.Model{ID: 2}, MenuName: "系统监控", ParentID: 0, OrderNum: 2, URL: "#", Target: "", MenuType: "M", Visible: "0", IsRefresh: "1", Perms: "", Icon: "fa fa-video-camera", Remark: "admin"},
		{Model: gorm.Model{ID: 3}, MenuName: "系统工具", ParentID: 0, OrderNum: 3, URL: "#", Target: "", MenuType: "M", Visible: "0", IsRefresh: "1", Perms: "", Icon: "fa fa-bars", Remark: "admin"},
		{Model: gorm.Model{ID: 4}, MenuName: "若依官网", ParentID: 0, OrderNum: 4, URL: "https://ruoyi.vip", Target: "menuBlank", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "", Icon: "fa fa-location-arrow", Remark: "admin"},
		{Model: gorm.Model{ID: 100}, MenuName: "用户管理", ParentID: 1, OrderNum: 1, URL: "/system/user", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "system:user:view", Icon: "fa fa-user-o", Remark: "admin"},
		{Model: gorm.Model{ID: 101}, MenuName: "角色管理", ParentID: 1, OrderNum: 2, URL: "/system/role", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "system:role:view", Icon: "fa fa-user-secret", Remark: "admin"},
		{Model: gorm.Model{ID: 102}, MenuName: "菜单管理", ParentID: 1, OrderNum: 3, URL: "/system/menu", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "system:menu:view", Icon: "fa fa-th-list", Remark: "admin"},
		{Model: gorm.Model{ID: 103}, MenuName: "部门管理", ParentID: 1, OrderNum: 4, URL: "/system/dept", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "system:dept:view", Icon: "fa fa-outdent", Remark: "admin"},
		{Model: gorm.Model{ID: 104}, MenuName: "岗位管理", ParentID: 1, OrderNum: 5, URL: "/system/post", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "system:post:view", Icon: "fa fa-address-card-o", Remark: "admin"},
		{Model: gorm.Model{ID: 105}, MenuName: "字典管理", ParentID: 1, OrderNum: 6, URL: "/system/dict", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "system:dict:view", Icon: "fa fa-bookmark-o", Remark: "admin"},
		{Model: gorm.Model{ID: 106}, MenuName: "参数设置", ParentID: 1, OrderNum: 7, URL: "/system/config", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "system:config:view", Icon: "fa fa-sun-o", Remark: "admin"},
		{Model: gorm.Model{ID: 107}, MenuName: "通知公告", ParentID: 1, OrderNum: 8, URL: "/system/notice", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "system:notice:view", Icon: "fa fa-bullhorn", Remark: "admin"},
		{Model: gorm.Model{ID: 108}, MenuName: "日志管理", ParentID: 1, OrderNum: 9, URL: "#", Target: "", MenuType: "M", Visible: "0", IsRefresh: "1", Perms: "", Icon: "fa fa-pencil-square-o", Remark: "admin"},
		{Model: gorm.Model{ID: 109}, MenuName: "在线用户", ParentID: 2, OrderNum: 1, URL: "/monitor/online", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "monitor:online:view", Icon: "fa fa-user-circle", Remark: "admin"},
		{Model: gorm.Model{ID: 110}, MenuName: "定时任务", ParentID: 2, OrderNum: 2, URL: "/monitor/job", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "monitor:job:view", Icon: "fa fa-tasks", Remark: "admin"},
		{Model: gorm.Model{ID: 111}, MenuName: "数据监控", ParentID: 2, OrderNum: 3, URL: "/monitor/data", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "monitor:data:view", Icon: "fa fa-bug", Remark: "admin"},
		{Model: gorm.Model{ID: 112}, MenuName: "服务监控", ParentID: 2, OrderNum: 4, URL: "/monitor/server", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "monitor:server:view", Icon: "fa fa-server", Remark: "admin"},
		{Model: gorm.Model{ID: 113}, MenuName: "缓存监控", ParentID: 2, OrderNum: 5, URL: "/monitor/cache", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "monitor:cache:view", Icon: "fa fa-cube", Remark: "admin"},
		{Model: gorm.Model{ID: 114}, MenuName: "表单构建", ParentID: 3, OrderNum: 1, URL: "/tool/build", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "tool:build:view", Icon: "fa fa-wpforms", Remark: "admin"},
		{Model: gorm.Model{ID: 115}, MenuName: "代码生成", ParentID: 3, OrderNum: 2, URL: "/tool/gen", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "tool:gen:view", Icon: "fa fa-code", Remark: "admin"},
		{Model: gorm.Model{ID: 116}, MenuName: "系统接口", ParentID: 3, OrderNum: 3, URL: "/tool/swagger", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "tool:swagger:view", Icon: "fa fa-gg", Remark: "admin"},
		{Model: gorm.Model{ID: 500}, MenuName: "操作日志", ParentID: 108, OrderNum: 1, URL: "/monitor/operlog", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "monitor:operlog:view", Icon: "fa fa-address-book", Remark: "admin"},
		{Model: gorm.Model{ID: 501}, MenuName: "登录日志", ParentID: 108, OrderNum: 2, URL: "/monitor/logininfor", Target: "", MenuType: "C", Visible: "0", IsRefresh: "1", Perms: "monitor:logininfor:view", Icon: "fa fa-file-image-o", Remark: "admin"},
	}
	return db.Save(menuData).Error
}
func userInit(db *gorm.DB) error {
	user := []model.SysUser{
		{DeptId: 103, UserType: "00", UserName: "admin", NickName: "超级管理员"},
		{DeptId: 105, UserType: "00", UserName: "user_one", NickName: "普通用户"},
	}
	pwd := "123456"
	user[0].Salt = utils.SaltSpawn()
	user[1].Salt = utils.SaltSpawn()
	user[0].Password = utils.EncryptPassword(user[0].UserName, pwd, user[0].Salt)
	user[1].Password = utils.EncryptPassword(user[1].UserName, pwd, user[1].Salt)
	return db.Save(user).Error
}
func bindInit(db *gorm.DB) error {
	roleMenu := []model.SysRole{
		{Model: gorm.Model{ID: 1}, Users: []model.SysUser{{Model: gorm.Model{ID: 1}}}},
		{
			Model: gorm.Model{ID: 2},
			Menus: []model.SysMenu{
				{Model: gorm.Model{ID: 1}},
				{Model: gorm.Model{ID: 2}},
				{Model: gorm.Model{ID: 3}},
				{Model: gorm.Model{ID: 4}},
				{Model: gorm.Model{ID: 100}},
				{Model: gorm.Model{ID: 101}},
				{Model: gorm.Model{ID: 102}},
				{Model: gorm.Model{ID: 103}},
				{Model: gorm.Model{ID: 104}},
				{Model: gorm.Model{ID: 105}},
				{Model: gorm.Model{ID: 106}},
				{Model: gorm.Model{ID: 107}},
				{Model: gorm.Model{ID: 108}},
				{Model: gorm.Model{ID: 109}},
				{Model: gorm.Model{ID: 110}},
				{Model: gorm.Model{ID: 111}},
				{Model: gorm.Model{ID: 112}},
				{Model: gorm.Model{ID: 113}},
				{Model: gorm.Model{ID: 114}},
				{Model: gorm.Model{ID: 115}},
				{Model: gorm.Model{ID: 116}},
				{Model: gorm.Model{ID: 500}},
				{Model: gorm.Model{ID: 501}},
			},
			Users: []model.SysUser{{Model: gorm.Model{ID: 2}}},
			Depts: []model.SysDept{{Model: gorm.Model{ID: 100}}, {Model: gorm.Model{ID: 101}}, {Model: gorm.Model{ID: 105}}},
		},
	}
	userDept := []model.SysUser{
		{Model: gorm.Model{ID: 1}, DeptId: 100},
		{Model: gorm.Model{ID: 2}, DeptId: 102},
	}
	err := db.Save(roleMenu).Error
	if err != nil {
		return err
	}
	err = db.Save(userDept).Error
	return err
}
