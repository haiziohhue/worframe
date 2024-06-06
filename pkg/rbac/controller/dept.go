package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"worframe/share/types"
)

type DeptController struct {
}

func (d *DeptController) GetAll(c *gin.Context) {
	q := types.NormalListQuery{}
	err := c.BindQuery(&q)
	if err != nil {
		panic(err)
	}
	c.Set("response_data", q)
	//c.JSON(http.StatusOK, q)
}
func (d *DeptController) GetOne(c *gin.Context) {
	//c.JSON(200, gin.H{})
	c.Set("response_data", nil)
}
func (d *DeptController) Update(c *gin.Context) {
	//c.JSON(200, gin.H{})
	c.Status(http.StatusOK)
}
func (d *DeptController) Delete(c *gin.Context) {

	c.Status(http.StatusOK)
	//c.JSON(200, gin.H{})
}
func (d *DeptController) Create(c *gin.Context) {
	//c.JSON(200, gin.H{})
	c.Status(http.StatusCreated)
	//c.Set("response_data", nil)
}
