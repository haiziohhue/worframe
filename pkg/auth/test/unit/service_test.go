package unit

import (
	"github.com/go-playground/assert/v2"
	"gorm.io/gorm"
	"testing"
	"worframe/pkg/auth/service"
	"worframe/share/model"
)

func TestDeptService(t *testing.T) {
	ds := service.NewDeptService(testApp.Logger, testApp.DB)
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
			DeptName:  "测试",
			Ancestors: "??",
		}
		err := ds.Create(dept)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, err, nil)
	})
	t.Run("Update", func(t *testing.T) {
		dept := &model.SysDept{
			Model:     gorm.Model{ID: 1},
			Ancestors: "???",
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
