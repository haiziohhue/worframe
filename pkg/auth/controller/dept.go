package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"worframe/pkg/auth/core/iface"
	"worframe/pkg/auth/service"
	"worframe/share/constant"
	"worframe/share/model"
	"worframe/share/types"
)

type DeptController struct {
	service service.DeptService
}

func NewDeptController(core iface.ICore) *DeptController {
	return &DeptController{
		service: *service.NewDeptService(core),
	}
}
func (ctrl *DeptController) GetAll(c *gin.Context) {
	q := types.NormalListQuery{}
	err := c.BindQuery(&q)
	if err != nil {
		_ = c.Error(err).SetType(constant.InvalidQuery)
		return
	}
	res, err := ctrl.service.FindAll(q.Page, q.PageSize)
	if err != nil {
		_ = c.Error(err).SetType(constant.RequestedResourceNotFound)
	}
	c.Set("response_data", res)
}
func (ctrl *DeptController) GetOne(c *gin.Context) {
	var q types.IdParam
	err := c.ShouldBindUri(&q)
	if err != nil {
		_ = c.Error(err).SetType(constant.InvalidQuery)
		return
	}
	res, err := ctrl.service.FindById(q.Id)
	if err != nil {
		_ = c.Error(err).SetType(constant.RequestedResourceNotFound)
		return
	}
	c.Set("response_data", res)
}
func (ctrl *DeptController) Update(c *gin.Context) {
	var dept model.SysDept
	err := c.BindJSON(&dept)
	if err != nil {
		_ = c.Error(err).SetType(constant.InvalidBody)
		return
	}
	var q types.IdParam
	err = c.ShouldBindUri(&q)
	if err != nil {
		_ = c.Error(err).SetType(constant.InvalidQuery)
		return
	}
	dept.ID = q.Id
	err = ctrl.service.Update(&dept)
	if err != nil {
		_ = c.Error(err).SetType(constant.RequestedResourceNotFound)
		return
	}
	c.Set("response_data", nil)
}
func (ctrl *DeptController) Delete(c *gin.Context) {
	var q types.IdParam
	err := c.ShouldBindUri(&q)
	if err != nil {
		_ = c.Error(err).SetType(constant.InvalidQuery)
		return
	}
	err = ctrl.service.Delete(q.Id)
	if err != nil {
		_ = c.Error(err).SetType(constant.RequestedResourceNotFound)
		return
	}
	c.Status(http.StatusOK)
}
func (ctrl *DeptController) Create(c *gin.Context) {
	var dept model.SysDept
	err := c.BindJSON(&dept)
	if err != nil {
		_ = c.Error(err).SetType(constant.InvalidBody)
		return
	}
	err = ctrl.service.Create(&dept)
	if err != nil {
		_ = c.Error(err).SetType(constant.RequestedResourceNotFound)
		return
	}
	c.Set("response_data", nil)
}
