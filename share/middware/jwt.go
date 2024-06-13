package middware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"worframe/share/config"
	"worframe/share/core/iface"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var (
	identityKey = "id"
	port        string
)

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}
type JwtPayload struct {
	UUID string   `json:"uuid"`
	Role []string `json:"role,omitempty"`
}

func JWTMiddleware(app iface.ICore) gin.HandlerFunc {
	//init
	middleware := jwt.New(initParams(*app.GetConf()))

	return func(c *gin.Context) {

	}
}

func handlerMiddleWare(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(context *gin.Context) {
		errInit := authMiddleware.MiddlewareInit()
		if errInit != nil {
			log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}
	}
}

// params init
func initParams(cfg config.Config) *jwt.GinJWTMiddleware {
	timeout, err := time.ParseDuration(cfg.Jwt.ExpiresTime)
	if err != nil {
		return nil
	}
	refresh, err := time.ParseDuration(cfg.Jwt.BufferTime)
	if err != nil {
		return nil
	}
	return &jwt.GinJWTMiddleware{
		Realm:       cfg.Server.Name,
		Key:         []byte(cfg.Jwt.SignKey),
		Timeout:     timeout,
		MaxRefresh:  refresh,
		IdentityKey: identityKey,
		// 载荷
		PayloadFunc: payloadFunc(cfg),
		//身份验证
		IdentityHandler: identityHandler(),
		//
		Authenticator: authenticator(),
		// 鉴权
		Authorizator: authorizator(),
		// 未鉴权成功处理
		Unauthorized: unauthorized(),

		//3种方式查找token
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}

func payloadFunc(cfg config.Config) func(data interface{}) jwt.MapClaims {
	expires, err := time.ParseDuration(cfg.ExpiresTime)
	if err != nil {
		expires = 5 * time.Minute
	}
	return func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*JwtPayload); ok {
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
		return &JwtPayload{
			Role: claims["role"].([]string),
			UUID: claims["uuid"].(string),
		}
	}
}

func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var loginVals login
		if err := c.ShouldBind(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		userID := loginVals.Username
		password := loginVals.Password

		if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
			return &User{
				UserName:  userID,
				LastName:  "Bo-Yi",
				FirstName: "Wu",
			}, nil
		}
		return nil, jwt.ErrFailedAuthentication
	}
}

func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		if v, ok := data.(*User); ok && v.UserName == "admin" {
			return true
		}
		return false
	}
}

func unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	}
}

func handleNoRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	}
}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*User).UserName,
		"text":     "Hello World.",
	})
}
