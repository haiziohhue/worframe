package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"worframe/pkg/auth/service"
	"worframe/share/constant"
	"worframe/share/model"
	"worframe/share/types"
)

type MenuController struct {
	service service.MenuService
}

func NewMenuController() *MenuController {
	return &MenuController{
		service: service.MenuService{},
	}
}
func (ctrl *MenuController) GetAll(c *gin.Context) {
	q := types.NormalListQuery{}
	err := c.BindQuery(&q)
	if err != nil {
		_ = c.Error(err).SetType(constant.INVALID_QUERY)
		return
	}
	res, err := ctrl.service.FindAll(q.Page, q.PageSize)
	if err != nil {
		_ = c.Error(err).SetType(constant.REQUESTED_RESOURCE_NOT_FOUND)
	}
	c.Set("response_data", res)
	c.Next()
}
func (ctrl *MenuController) GetOne(c *gin.Context) {
	var q types.IdParam
	err := c.ShouldBindUri(&q)
	if err != nil {
		_ = c.Error(err).SetType(constant.INVALID_QUERY)
		return
	}
	res, err := ctrl.service.FindById(q.Id)
	if err != nil {
		_ = c.Error(err).SetType(constant.REQUESTED_RESOURCE_NOT_FOUND)
		return
	}
	c.Set("response_data", res)
}
func (ctrl *MenuController) Update(c *gin.Context) {
	var menu model.SysMenu
	err := c.BindJSON(&menu)
	if err != nil {
		_ = c.Error(err).SetType(constant.INVALID_BODY)
		return
	}
	var q types.IdParam
	err = c.ShouldBindUri(&q)
	if err != nil {
		_ = c.Error(err).SetType(constant.INVALID_QUERY)
		return
	}
	menu.ID = q.Id
	err = ctrl.service.Update(&menu)
	if err != nil {
		_ = c.Error(err).SetType(constant.REQUESTED_RESOURCE_NOT_FOUND)
		return
	}
	c.Set("response_data", nil)
}
func (ctrl *MenuController) Delete(c *gin.Context) {
	var q types.IdParam
	err := c.ShouldBindUri(&q)
	if err != nil {
		_ = c.Error(err).SetType(constant.INVALID_QUERY)
		return
	}
	err = ctrl.service.Delete(q.Id)
	if err != nil {
		_ = c.Error(err).SetType(constant.REQUESTED_RESOURCE_NOT_FOUND)
		return
	}
	c.Status(http.StatusOK)
}
func (ctrl *MenuController) Create(c *gin.Context) {
	var menu model.SysMenu
	err := c.BindJSON(&menu)
	if err != nil {
		_ = c.Error(err).SetType(constant.INVALID_BODY)
		return
	}
	err = ctrl.service.Create(&menu)
	if err != nil {
		_ = c.Error(err).SetType(constant.REQUESTED_RESOURCE_NOT_FOUND)
		return
	}
	c.Set("response_data", nil)
}
