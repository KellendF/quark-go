package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/quarkcms/quark-go/pkg/framework/config"
	"github.com/quarkcms/quark-go/pkg/framework/rand"
)

var appKey string

// 初始化
func init() {
	key := config.Get("app.key").(string)
	if key == "" {
		key = rand.Make("alphanumeric", 950)
	}
	appKey = key
}

// 创建token
func Make(mapClaims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return token.SignedString([]byte(appKey))
}

// 解析token
func Parse(token string) (jwt.MapClaims, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(appKey), nil
	})

	if err != nil {
		return nil, err
	}

	return claim.Claims.(jwt.MapClaims), nil
}
