package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type ApiMiddleware struct{}

// 接口中间件
func (p *ApiMiddleware) Handle(c *fiber.Ctx) error {

	return c.Next()
}
