package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/framework/token"
)

type AdminMiddleware struct{}

// 后台中间件
func (p *AdminMiddleware) Handle(c *fiber.Ctx) error {

	// token认证
	header := c.GetReqHeaders()
	getToken := strings.Split(header["Authorization"], " ")

	if len(getToken) != 2 {
		return c.SendStatus(401)
	}

	userInfo, err := token.Parse(getToken[1])
	if err != nil {
		fmt.Println(err)
	}

	if userInfo["guard_name"] != "admin" {
		return c.SendStatus(401)
	}

	return c.Next()
}
