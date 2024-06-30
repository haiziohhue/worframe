package unit

import (
	"github.com/go-playground/assert/v2"
	"gorm.io/gorm"
	"testing"
	"worframe/pkg/auth/migrate"
	"worframe/share/model"
)

func TestAuthService(t *testing.T) {
	mi := migrate.NewDBMigrate(testApp.GetDB())
	err := mi.TestEnvInit(testApp.GetLog())
	if err != nil {
		panic(err)
	}
	t.Run("dept", testDeptService)
	t.Run("menu", testMenuService)
	t.Run("role", testRoleService)

}
func testDeptService(t *testing.T) {

	ds := application.NewDeptService(testApp)
	t.Run("GetAll", func(t *testing.T) {
		resp, err := ds.FindAll(1, 10)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(&resp)
	})

	t.Run("GetOne", func(t *testing.T) {
		resp, err := ds.FindById(1)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp)
		assert.Equal(t, resp.ID, uint(1))
	})
	t.Run("Create", func(t *testing.T) {
		dept := &model.SysDept{
			DeptName: "测试Create",
		}
		err := ds.Create(dept)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, err, nil)
	})
	t.Run("Update", func(t *testing.T) {
		dept := &model.SysDept{
			Model:    gorm.Model{ID: 1},
			DeptName: "测试Update",
		}
		err := ds.Update(dept)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, err, nil)
	})
	t.Run("Delete", func(t *testing.T) {
		err := ds.Delete(1)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, err, nil)
	})
}
func testMenuService(t *testing.T) {

	ds := application.NewMenuService(testApp)
	t.Run("GetAll", func(t *testing.T) {
		resp, err := ds.FindAll(1, 10)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(&resp)
	})

	t.Run("GetOne", func(t *testing.T) {
		resp, err := ds.FindById(1)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp)
		assert.Equal(t, resp.ID, uint(1))
	})
	t.Run("Create", func(t *testing.T) {
		menu := &model.SysMenu{
			MenuName: "testMenuName",
		}
		err := ds.Create(menu)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, err, nil)
	})
	t.Run("Update", func(t *testing.T) {
		menu := &model.SysMenu{
			Model:    gorm.Model{ID: 1},
			MenuName: "testService修改权限名",
		}
		err := ds.Update(menu)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, err, nil)
	})
	t.Run("Delete", func(t *testing.T) {
		err := ds.Delete(1)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, err, nil)
	})
}
func testRoleService(t *testing.T) {

	ds := application.NewRoleService(testApp)
	t.Run("GetAll", func(t *testing.T) {
		resp, err := ds.FindAll(1, 10)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(&resp)
	})

	t.Run("GetOne", func(t *testing.T) {
		resp, err := ds.FindById(1)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp)
		assert.Equal(t, resp.ID, uint(1))
	})
	t.Run("Create", func(t *testing.T) {
		role := &model.SysRole{
			RoleName: "testRoleName",
			RoleSort: 1,
		}
		err := ds.Create(role)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, err, nil)
	})
	t.Run("Update", func(t *testing.T) {
		role := &model.SysRole{
			Model:   gorm.Model{ID: 1},
			RoleKey: "测试Key",
		}
		err := ds.Update(role)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, err, nil)
	})
	t.Run("Delete", func(t *testing.T) {
		err := ds.Delete(6)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, err, nil)
	})
}
