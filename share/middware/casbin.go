package middware

import (
	"github.com/casbin/casbin/util"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	redisadapter "github.com/casbin/redis-adapter/v3"
	"github.com/gin-gonic/gin"
	"net/http"
	"worframe/pkg/auth/utils"
	"worframe/share/constant"
	"worframe/share/core"
)

func Casbin() gin.HandlerFunc {
	m := modelBind(core.Cfg.Casbin.ModelName)
	redisAdapt, err := redisadapter.NewAdapterWithPool(core.Redis)
	if err != nil {
		panic(err)
	}
	e, _ := casbin.NewEnforcer(m, redisAdapt)
	e.AddNamedMatchingFunc("g", "KeyMatch2", util.KeyMatch2)
	a := &BasicAuthorizer{enforcer: e}
	return func(c *gin.Context) {
		if !a.CheckPermission(c) {
			a.RequirePermission(c)
		}
	}
}

/*
modelBind 绑定模型
*/
func modelBind(name string) model.Model {
	ms := constant.CasbinModel[name]
	if ms == "" {
		panic("no found model" + name)
	}
	m, err := model.NewModelFromString(ms)
	if err != nil {
		panic(err)
	}
	return m
}

type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

func (a *BasicAuthorizer) CheckPermission(c *gin.Context) bool {
	perm := utils.GetUserPerm(c)
	method := c.Request.Method
	path := c.Request.URL.Path
	for _, sub := range perm {
		allowed, err := a.enforcer.Enforce(sub, path, method)
		if err == nil && allowed {
			return true
		}
	}
	return false
}

// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}
