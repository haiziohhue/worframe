package controller

import "github.com/gin-gonic/gin"

type RoleController struct {
}

func (ctrl *RoleController) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{})
}
func (ctrl *RoleController) GetOne(c *gin.Context) {
	c.JSON(200, gin.H{})
}
func (ctrl *RoleController) Update(c *gin.Context) {
	c.JSON(200, gin.H{})
}
func (ctrl *RoleController) Delete(c *gin.Context) {
	c.JSON(200, gin.H{})
}
func (ctrl *RoleController) Create(c *gin.Context) {
	c.JSON(200, gin.H{})
}
