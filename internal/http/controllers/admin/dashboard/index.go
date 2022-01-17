package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/component/page"
)

// 主页面
func Index(c *fiber.Ctx) error {

	pageComponent := &page.Component{}

	component := pageComponent.
		SetTitle("测试").
		SetBody("test").
		JsonSerialize()

	return c.JSON(component)
}
