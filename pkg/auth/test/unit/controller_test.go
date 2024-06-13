package unit

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"worframe/pkg/auth/migrate"
	"worframe/share/model"
	"worframe/share/types"
)

type header struct {
	Key   string
	Value string
}

func performRequest(r http.Handler, method, path string, body []byte, headers ...header) *httptest.ResponseRecorder {
	data := bytes.NewReader(body)
	req := httptest.NewRequest(method, path, data)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestAuthMethod(t *testing.T) {
	mi := migrate.NewDBMigrate(testApp.GetDB())
	err := mi.TestEnvInit(testApp.GetLog())
	if err != nil {
		panic(err)
	}
	t.Run("Dept", testDeptMethod)
	t.Run("Dept", testRoleMethod)
	t.Run("Dept", testMenuMethod)
}
func testDeptMethod(t *testing.T) {
	t.Run("GetOne", func(t *testing.T) {
		w := performRequest(testApp.GetEngine(), http.MethodGet, "/dept/1", nil)
		var res types.BaseRes[model.SysDept]
		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Data.ID, uint(1))
	})
	t.Run("GetOneNotFound", func(t *testing.T) {
		w := performRequest(testApp.GetEngine(), http.MethodGet, "/dept/200", nil)
		var res types.BaseRes[model.SysDept]
		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Code, 40009)
	})
	t.Run("Create", func(t *testing.T) {
		dept := model.SysDept{
			ParentID:  0,
			Ancestors: "0",
			DeptName:  "项目总部",
			OrderNum:  0,
			Leader:    "海石花",
			Phone:     "13700000000",
			Email:     "test@test.com",
		}
		deptBytes, _ := json.Marshal(dept)
		w := performRequest(testApp.GetEngine(), http.MethodPost, "/dept", deptBytes)
		var res types.EmptyRes

		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Code, 20000)
	})
	t.Run("Update", func(t *testing.T) {
		dept := model.SysDept{
			DeptName: "测试更新部门名称100",
		}
		deptBytes, _ := json.Marshal(dept)

		w := performRequest(testApp.GetEngine(), http.MethodPut, "/dept/3", deptBytes)
		var res types.EmptyRes
		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Code, 20000)
	})
	t.Run("Delete", func(t *testing.T) {
		dept := model.SysDept{
			DeptName: "测试更新部门名称100",
		}
		deptBytes, _ := json.Marshal(dept)

		w := performRequest(testApp.GetEngine(), http.MethodDelete, "/dept/4", deptBytes)
		var res types.EmptyRes
		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Code, 20000)
	})
}
func testRoleMethod(t *testing.T) {
	t.Run("GetOne", func(t *testing.T) {
		w := performRequest(testApp.GetEngine(), http.MethodGet, "/role/1", nil)
		var res types.BaseRes[model.SysRole]
		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Data.ID, uint(1))
	})
	t.Run("GetOneNotFound", func(t *testing.T) {
		w := performRequest(testApp.GetEngine(), http.MethodGet, "/role/200", nil)
		var res types.BaseRes[model.SysRole]
		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Code, 40009)
	})
	t.Run("Create", func(t *testing.T) {
		role := model.SysRole{
			RoleName: "测试创建新权限",
		}
		roleBytes, _ := json.Marshal(role)
		w := performRequest(testApp.GetEngine(), http.MethodPost, "/role", roleBytes)
		var res types.EmptyRes

		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Code, 20000)
	})
	t.Run("Update", func(t *testing.T) {
		role := model.SysRole{
			RoleName: "测试更新Update",
			Depts: []*model.SysDept{
				{DeptName: "roleUpdate测试更新部门"},
			},
		}
		roleBytes, _ := json.Marshal(role)

		w := performRequest(testApp.GetEngine(), http.MethodPut, "/role/1", roleBytes)
		var res types.EmptyRes
		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Code, 20000)
	})
	t.Run("Delete", func(t *testing.T) {
		w := performRequest(testApp.GetEngine(), http.MethodDelete, "/role/4", nil)
		var res types.EmptyRes
		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Code, 20000)
	})
}
func testMenuMethod(t *testing.T) {
	t.Run("GetOne", func(t *testing.T) {
		w := performRequest(testApp.GetEngine(), http.MethodGet, "/menu/1", nil)
		var res types.BaseRes[model.SysMenu]
		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Data.ID, uint(1))
	})
	t.Run("GetOneNotFound", func(t *testing.T) {
		w := performRequest(testApp.GetEngine(), http.MethodGet, "/menu/200", nil)
		var res types.BaseRes[model.SysMenu]
		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Code, 40009)
	})
	t.Run("Create", func(t *testing.T) {
		menu := model.SysMenu{
			ParentID: 0,
			OrderNum: 0,
			MenuName: "测试CreatePerm",
			MenuType: "1234",
			Method:   "GET",
			Roles: []*model.SysRole{
				{Model: gorm.Model{ID: 1}},
			},
		}
		menuBytes, _ := json.Marshal(menu)
		w := performRequest(testApp.GetEngine(), http.MethodPost, "/menu", menuBytes)
		var res types.EmptyRes

		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Code, 20000)
	})
	t.Run("Update", func(t *testing.T) {
		menu := model.SysDept{
			DeptName: "测试更新部门名称100和绑定权限5",
			Roles:    []*model.SysRole{{Model: gorm.Model{ID: 5}}},
		}
		menuBytes, _ := json.Marshal(menu)

		w := performRequest(testApp.GetEngine(), http.MethodPut, "/menu/3", menuBytes)
		var res types.EmptyRes
		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Code, 20000)
	})
	t.Run("Delete", func(t *testing.T) {

		w := performRequest(testApp.GetEngine(), http.MethodDelete, "/menu/4", nil)
		var res types.EmptyRes
		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Code, 20000)
	})
}
