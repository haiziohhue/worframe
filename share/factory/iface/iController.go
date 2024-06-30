package iface

import "github.com/gin-gonic/gin"

type IController interface {
	// GetOne 获取单个数据
	GetOne() gin.HandlerFunc
	// GetAll 获取所有数据
	GetAll() gin.HandlerFunc
	// Create 保存数据
	Create() gin.HandlerFunc
	// Update 更新数据
	Update() gin.HandlerFunc
	// Delete 删除数据
	Delete() gin.HandlerFunc
}
