package unit

import (
	"github.com/go-playground/assert/v2"
	"gorm.io/gorm"
	"testing"
	"worframe/pkg/rbac/service"
	"worframe/share/model"
)

func TestDeptServiceGetAll(t *testing.T) {
	ds := service.DeptService{}
	resp, err := ds.FindAll(1, 10)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(&resp)
}
func TestDeptServiceGetOne(t *testing.T) {
	ds := service.DeptService{}
	resp, err := ds.FindById(101)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
	assert.Equal(t, resp.ID, uint(101))
}
func TestDeptServiceCreate(t *testing.T) {
	ds := service.DeptService{}
	dept := &model.SysDept{
		DeptName:  "测试",
		Ancestors: "??",
	}
	err := ds.Create(dept)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, err, nil)
}
func TestDeptServiceUpdate(t *testing.T) {
	ds := service.DeptService{}
	dept := &model.SysDept{
		Model:     gorm.Model{ID: 101},
		Ancestors: "???",
	}
	err := ds.Update(dept)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, err, nil)
}
func TestDeptServiceDelete(t *testing.T) {
	ds := service.DeptService{}
	err := ds.Delete(101)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, err, nil)
}
