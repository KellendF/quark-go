package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/sessions"
)

// 全局中间件
func App(c *fiber.Ctx) error {
	// 初始化session
	sessions.Init(c)

	return c.Next()
}
