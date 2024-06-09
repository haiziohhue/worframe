package utils

import (
	"errors"
	"github.com/gofrs/uuid/v5"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/sync/singleflight"
	"time"
	"worframe/pkg/auth/config"
	"worframe/share/core"
	"worframe/share/utils"
)

type JWT struct {
}
type BaseClaims struct {
	UUID     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Dept     string    `json:"dept"`
	Roles    []string  `json:"roles"`
	Menus    []string  `json:"menus"`
}
type CustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims
	BufferTime int64
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	bf, _ := utils.ParseDuration(config.AuthCfg.BufferTime)
	ep, _ := utils.ParseDuration(config.AuthCfg.ExpiresTime)
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间 7天  配置文件
			Issuer:    core.Cfg.Name,                             // 签名的发行者
		},
	}
	return claims
}

func (j *JWT) CreateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.AuthCfg.SignKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
	var controller = singleflight.Group{}
	v, err, _ := controller.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return config.AuthCfg.SignKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
