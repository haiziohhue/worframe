package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/util"
	redisadapter "github.com/casbin/redis-adapter/v3"
	"github.com/gin-gonic/gin"
	"time"
	iface2 "worframe/pkg/auth/core/iface"
	"worframe/share/config"
	"worframe/share/constant"
	"worframe/share/core/iface"
	"worframe/share/types"
)

type casbinAuthorizer struct {
	enforcer *casbin.Enforcer
}

func JWTMiddleware(app *iface.ICore) gin.HandlerFunc {
	middleware, err := jwt.New(initParams(app))
	if err != nil {
		(*app).SetErr(err)
	}
	return middleware.MiddlewareFunc()
}

// NewAuthRouter auth package use it
func NewAuthRouter(app iface2.ICore, loginFn func(c *gin.Context) (interface{}, error)) (*jwt.GinJWTMiddleware, error) {
	jwtMiddleware := initParams(app.GetRawCore())
	jwtMiddleware.Authenticator = loginFn
	return jwt.New(jwtMiddleware)
}

func initParams(app *iface.ICore) *jwt.GinJWTMiddleware {

	timeout, err := time.ParseDuration((*app).GetConf().Jwt.ExpiresTime)
	if err != nil {
		return nil
	}
	refresh, err := time.ParseDuration((*app).GetConf().Jwt.BufferTime)
	if err != nil {
		return nil
	}
	return &jwt.GinJWTMiddleware{
		Realm:       (*app).GetConf().Server.Name,
		Key:         []byte((*app).GetConf().Jwt.SignKey),
		Timeout:     timeout,
		MaxRefresh:  refresh,
		IdentityKey: "AuthPassRole",

		PayloadFunc:     payloadFunc((*app).GetConf()),
		IdentityHandler: identityHandler(),
		Authorizator:    authorizator(app),
		Unauthorized:    unauthorized(),

		//3种方式查找token
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}

func payloadFunc(cfg *config.Config) func(data interface{}) jwt.MapClaims {
	expires, err := time.ParseDuration(cfg.ExpiresTime)
	if err != nil {
		expires = 5 * time.Minute
	}
	return func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*types.JwtPayload); ok {
			claims := jwt.MapClaims{}
			claims["iat"] = time.Now().Unix()
			claims["iss"] = cfg.Server.Name
			claims["exp"] = time.Now().Add(expires).Unix()
			claims["sub"] = "User"
			//claims["aud"] = "audience"
			claims["uuid"] = v.UUID
			claims["role"] = v.Role
			return claims
		}
		return jwt.MapClaims{}
	}
}

func identityHandler() func(c *gin.Context) any {
	return func(c *gin.Context) any {
		claims := jwt.ExtractClaims(c)
		return &types.JwtPayload{
			Role: claims["role"].([]string),
			UUID: claims["uuid"].(string),
		}
	}
}

func authorizator(app *iface.ICore) func(data interface{}, c *gin.Context) bool {
	m := modelBind((*app).GetConf().Casbin.ModelName)
	redisAdapt, err := redisadapter.NewAdapterWithPool((*app).GetRedis())
	if err != nil {
		panic(err)
	}
	e, _ := casbin.NewEnforcer(m, redisAdapt)
	e.AddNamedMatchingFunc("g", "KeyMatch2", util.KeyMatch2)
	a := &casbinAuthorizer{enforcer: e}
	return func(data interface{}, c *gin.Context) bool {
		if v, ok := data.(*types.JwtPayload); ok {
			v.Role = append(v.Role, v.UUID)
			if !a.CheckPermission(v.Role, c) {
				return false
			}
			return true
		} else {
			return false
		}
	}
}

// 直接返回未验证
func unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	}
}

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

func (a *casbinAuthorizer) CheckPermission(perms []string, c *gin.Context) bool {
	method := c.Request.Method
	path := c.Request.URL.Path
	for _, sub := range perms {
		allowed, err := a.enforcer.Enforce(sub, path, method)
		if err == nil && allowed {
			return true
		}
	}
	return false
}
