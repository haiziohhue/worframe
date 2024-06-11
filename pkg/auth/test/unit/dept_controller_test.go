package unit

import (
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"net/http"
	"testing"
	"worframe/pkg/auth/migrate"
	"worframe/share/model"
	"worframe/share/types"
)

func TestDeptMethod(t *testing.T) {

	mi := migrate.NewDBMigrate(testApp.DB)
	err := mi.TestEnvInit(testApp.Logger)
	if err != nil {
		panic(err)
	}

	t.Run("deptGetOne", func(t *testing.T) {
		w := performRequest(testApp.Engine, http.MethodGet, "/dept/1", nil)
		var res types.BaseRes[model.SysDept]
		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Data.ID, uint(1))
	})
	t.Run("deptGetOneNotFound", func(t *testing.T) {
		w := performRequest(testApp.Engine, http.MethodGet, "/dept/200", nil)
		var res types.BaseRes[model.SysDept]
		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Code, 40009)
	})
	t.Run("deptCreate", func(t *testing.T) {
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
		w := performRequest(testApp.Engine, http.MethodPost, "/dept", deptBytes)
		var res types.EmptyRes

		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Code, 20000)
	})
	t.Run("deptUpdate", func(t *testing.T) {
		dept := model.SysDept{
			DeptName: "测试更新部门名称100",
		}
		deptBytes, _ := json.Marshal(dept)

		w := performRequest(testApp.Engine, http.MethodPut, "/dept/3", deptBytes)
		var res types.EmptyRes
		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Code, 20000)
	})
	t.Run("deptDelete", func(t *testing.T) {
		dept := model.SysDept{
			DeptName: "测试更新部门名称100",
		}
		deptBytes, _ := json.Marshal(dept)

		w := performRequest(testApp.Engine, http.MethodDelete, "/dept/4", deptBytes)
		var res types.EmptyRes
		err := json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Code, 20000)
	})
}
