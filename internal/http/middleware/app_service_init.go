package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type AppServiceInit struct{}

// 全局中间件
func (p *AppServiceInit) Handle(c *fiber.Ctx) error {
	return c.Next()
}
