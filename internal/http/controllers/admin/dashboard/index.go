package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/component/layout"
	"github.com/quarkcms/quark-go/pkg/component/page"
)

// 主页面
func Index(c *fiber.Ctx) error {

	layoutComponent := &layout.Component{}

	pageComponent := &page.Component{}

	getComponent := pageComponent.
		SetTitle("测试").
		SetBody("test").
		JsonSerialize()

	component := layoutComponent.SetTitle("test").
		SetBody(getComponent).
		JsonSerialize()

	return c.JSON(component)
}
