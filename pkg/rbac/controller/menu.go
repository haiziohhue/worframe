package controller

import "github.com/gin-gonic/gin"

type MenuController struct {
}

func (ctrl *MenuController) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{})
}
func (ctrl *MenuController) GetOne(c *gin.Context) {
	c.JSON(200, gin.H{})
}
func (ctrl *MenuController) Update(c *gin.Context) {
	c.JSON(200, gin.H{})
}
func (ctrl *MenuController) Delete(c *gin.Context) {
	c.JSON(200, gin.H{})
}
func (ctrl *MenuController) Create(c *gin.Context) {
	c.JSON(200, gin.H{})
}
