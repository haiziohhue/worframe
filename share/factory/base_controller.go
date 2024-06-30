package factory

import (
	"github.com/gin-gonic/gin"
	"worframe/share/constant"
	"worframe/share/factory/iface"
)

type BaseController[T iface.IDto, E iface.IEntity, D iface.IDao] struct {
	S iface.IWebService[T, E, D]
}
type ListQuery struct {
	Page     int `form:"page,default=1" json:"page"`
	PageSize int `form:"page_size,default=10" json:"page_size"`
}
type IDParam struct {
	ID uint `url:"id"`
}

func (ctrl BaseController[T, E, D]) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		var Param IDParam
		if err := c.ShouldBindUri(&Param); err != nil {
			c.Error(err)
		}
		dto, err := ctrl.S.GetOne(Param.ID)
		if err != nil {
			c.Error(err)
		}
		c.Set(constant.Response, dto)
	}
}

func (ctrl BaseController[T, E, D]) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		var p ListQuery
		if err := c.ShouldBindQuery(&p); err != nil {
			c.Error(err)
		}
		list, err := ctrl.S.GetAll(p.Page, p.PageSize)
		if err != nil {
			c.Error(err)
		}
		c.Set(constant.Response, list)
	}
}

func (ctrl BaseController[T, E, D]) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto T
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.Error(err)
		}
		if err := ctrl.S.Create(dto); err != nil {
			c.Error(err)
		}
	}
}

func (ctrl BaseController[T, E, D]) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto T
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.Error(err)
		}
		var Param IDParam
		if err := c.ShouldBindUri(&Param); err != nil {
			c.Error(err)
		}
		if err := ctrl.S.Update(Param.ID, dto); err != nil {
			c.Error(err)
		}
	}
}

func (ctrl BaseController[T, E, D]) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var Param IDParam
		if err := c.ShouldBindUri(&Param); err != nil {
			c.Error(err)
		}
		if err := ctrl.S.Delete(Param.ID); err != nil {
			c.Error(err)
		}
	}
}

func NewBaseController[T iface.IDto, E iface.IEntity, D iface.IDao](service iface.IWebService[T, E, D]) iface.IController {
	return BaseController[T, E, D]{
		S: service,
	}
}
