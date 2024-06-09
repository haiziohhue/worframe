package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"net"
)

func ClearToken(c *gin.Context) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("s-token", "", -1, "/", "", false, false)
	} else {
		c.SetCookie("s-token", "", -1, "/", host, false, false)
	}
}

func SetToken(c *gin.Context, token string, maxAge int) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("s-token", token, maxAge, "/", "", false, false)
	} else {
		c.SetCookie("s-token", token, maxAge, "/", host, false, false)
	}
}

func GetToken(c *gin.Context) string {
	token, _ := c.Cookie("s-token")
	if token == "" {
		token = c.Request.Header.Get("s-token")
	}
	return token
}

func GetClaims(c *gin.Context) (*CustomClaims, error) {
	token := GetToken(c)
	j := &JWT{}
	claims, err := j.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return claims, err
}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.UUID
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse
	}
}

// GetUserName 从Gin的Context中获取从jwt解析出来的用户名
func GetUserName(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.Username
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.Username
	}
}
func GetUserPerm(c *gin.Context) []string {
	if _, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return []string{}
		} else {
			perms := make([]string, 0)
			if cl.Menus != nil {
				perms = append(perms, cl.Menus...)
			}
			if cl.Roles != nil {
				perms = append(perms, cl.Roles...)
			}
			if cl.Dept != "" {
				perms = append(perms, cl.Dept)
			}
			return perms
		}
	}
	return []string{}
}
