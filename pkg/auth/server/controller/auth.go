package controller

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"worframe/pkg/auth/application"
	"worframe/pkg/auth/core/iface"
	"worframe/share/constant"
)

type AuthController struct {
	AuthService application.AuthService
}

func NewAuthController(core iface.ICore) *AuthController {
	Service, err := application.NewAuthService(&core)
	if err != nil {
		panic(err)
	}
	return &AuthController{
		AuthService: *Service,
	}
}

func (ctrl *AuthController) LoginPkg() func(c *gin.Context) (interface{}, error) {
	C := ctrl
	return func(c *gin.Context) (interface{}, error) {

		method, ok := c.GetQuery("method")
		if !ok {
			err := fmt.Errorf("invalid request")
			_ = c.Error(err).SetType(constant.InvalidQuery)
			return nil, err
		}
		j := application.LoginParams{}
		err := c.ShouldBindJSON(&j)
		if err != nil {
			_ = c.Error(err).SetType(constant.InvalidBody)
		}
		jwtPayload, err := C.AuthService.Login(j, method)
		if err != nil {
			return nil, err
		}
		return jwtPayload, err
	}
}
func (ctrl *AuthController) Register(c *gin.Context) {
	param := application.LoginParams{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		_ = c.Error(err).SetType(constant.InvalidBody)
		return
	}
	method, ok := c.GetQuery("method")
	if !ok {
		method = "password"
	}
	err = ctrl.AuthService.Register(param, method)
	if err != nil {
		_ = c.Error(err).SetType(constant.InvalidQuery)
		return
	}
	c.Status(http.StatusOK)
}
