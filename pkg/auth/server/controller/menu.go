package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"worframe/pkg/auth/core/iface"
	"worframe/share/constant"
	"worframe/share/model"
	"worframe/share/types"
)

type MenuController struct {
	service *application.MenuService
}

func NewMenuController(core iface.ICore) *MenuController {
	return &MenuController{
		service: application.NewMenuService(core),
	}
}
func (ctrl *MenuController) GetAll(c *gin.Context) {
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
func (ctrl *MenuController) GetOne(c *gin.Context) {
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
func (ctrl *MenuController) Update(c *gin.Context) {
	var menu model.SysMenu
	err := c.BindJSON(&menu)
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
	menu.ID = q.Id
	err = ctrl.service.Update(&menu)
	if err != nil {
		_ = c.Error(err).SetType(constant.RequestedResourceNotFound)
		return
	}
	c.Set("response_data", nil)
}
func (ctrl *MenuController) Delete(c *gin.Context) {
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
func (ctrl *MenuController) Create(c *gin.Context) {
	var menu model.SysMenu
	err := c.BindJSON(&menu)
	if err != nil {
		_ = c.Error(err).SetType(constant.InvalidBody)
		return
	}
	err = ctrl.service.Create(&menu)
	if err != nil {
		_ = c.Error(err).SetType(constant.RequestedResourceNotFound)
		return
	}
	c.Set("response_data", nil)
}
