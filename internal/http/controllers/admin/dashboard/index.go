package dashboard

import (
	"github.com/gofiber/fiber/v2"
)

// 表结构体
type Component struct {
	Component string `json:"component"`
	Body      string `json:"body"`
}

// 主页面
func Index(c *fiber.Ctx) error {

	Component := &Component{
		Component: "page",
		Body:      "test",
	}

	return c.JSON(Component)
}
