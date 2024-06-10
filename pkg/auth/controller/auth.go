package controller

import (
	"github.com/gin-gonic/gin"
	"worframe/share/constant"
	"worframe/share/types"
)

type AuthController struct {
}

func (ctrl *AuthController) LoginByPassword(c *gin.Context) {
	var data types.LoginReqBody
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.Error(err).SetType(constant.InvalidBody)
	}

}
func (ctrl *AuthController) LoginByPhone(c *gin.Context) {

}
func (ctrl *AuthController) LoginByEmail(c *gin.Context) {

}
func (ctrl *AuthController) Logout(c *gin.Context) {

}
func (ctrl *AuthController) Register(c *gin.Context) {

}
func (ctrl *AuthController) ForgotPassword(c *gin.Context) {

}
func (ctrl *AuthController) SetPassword(c *gin.Context) {

}
