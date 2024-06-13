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

type RoleController struct {
	service *service.RoleService
}

func NewRoleController(core iface.ICore) *RoleController {
	return &RoleController{
		service: service.NewRoleService(core),
	}
}

func (ctrl *RoleController) GetAll(c *gin.Context) {
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
	c.Next()
}
func (ctrl *RoleController) GetOne(c *gin.Context) {
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
func (ctrl *RoleController) Update(c *gin.Context) {
	var role model.SysRole
	err := c.BindJSON(&role)
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
	role.ID = q.Id
	err = ctrl.service.Update(&role)
	if err != nil {
		_ = c.Error(err).SetType(constant.RequestedResourceNotFound)
		return
	}
	c.Set("response_data", nil)
}
func (ctrl *RoleController) Delete(c *gin.Context) {
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
func (ctrl *RoleController) Create(c *gin.Context) {
	var role model.SysRole
	err := c.BindJSON(&role)
	if err != nil {
		_ = c.Error(err).SetType(constant.InvalidBody)
		return
	}
	err = ctrl.service.Create(&role)
	if err != nil {
		_ = c.Error(err).SetType(constant.RequestedResourceNotFound)
		return
	}
	c.Set("response_data", nil)
}
