package gin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"testing"
	"worframe/pkg/auth/migrate"
	"worframe/pkg/auth/server"
	"worframe/share/core"
	"worframe/share/initialize"
	"worframe/share/model"
	"worframe/share/types"
)

func TestDeptMethod(t *testing.T) {
	core.Cfg = initialize.InitConfig("test")
	core.DB = initialize.InitGorm(core.Cfg)
	m := migrate.NewDBMigrate(core.DB)
	err := m.TestEnvInit()
	if err != nil {
		t.Fatal(err)
	}
	r := gin.New()
	server.AuthInitServer(r)
	t.Run("deptGetOne", func(t *testing.T) {
		w := performRequest(r, http.MethodGet, "/dept/100", nil)
		var res types.BaseRes[model.SysDept]
		err = json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Data.ID, uint(100))
	})
	t.Run("deptGetOneNotFound", func(t *testing.T) {
		w := performRequest(r, http.MethodGet, "/dept/200", nil)
		var res types.BaseRes[model.SysDept]
		err = json.Unmarshal(w.Body.Bytes(), &res)
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
		w := performRequest(r, http.MethodPost, "/dept", deptBytes)
		var res types.EmptyRes

		err = json.Unmarshal(w.Body.Bytes(), &res)
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

		w := performRequest(r, http.MethodPut, "/dept/101", deptBytes)
		var res types.EmptyRes
		err = json.Unmarshal(w.Body.Bytes(), &res)
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

		w := performRequest(r, http.MethodDelete, "/dept/102", deptBytes)
		var res types.EmptyRes
		err = json.Unmarshal(w.Body.Bytes(), &res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, res.Code, 20000)
	})
	//w2 := performRequest(r, http.MethodGet, "/dept/100")
	//w3 := performRequest(r, http.MethodDelete, "/dept/109")
	//assert.Equal(t, http.StatusOK, w2.Code)
	//assert.Equal(t, http.StatusOK, w3.Code)
}
