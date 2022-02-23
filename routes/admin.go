package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/admin/controllers"
	"github.com/quarkcms/quark-go/internal/admin/dashboards"
	"github.com/quarkcms/quark-go/internal/admin/middleware"
	"github.com/quarkcms/quark-go/internal/admin/resources"
	"github.com/quarkcms/quark-go/pkg/ui/admin/dashboard"
	"github.com/quarkcms/quark-go/pkg/ui/admin/resource"
)

type Admin struct{}

// 仪表盘
var dashboardService = map[string]dashboard.DashboardInterface{
	"index": &dashboards.Index{},
}

// 资源
var resourceService = map[string]resource.ResourceInterface{
	"admin": &resources.Admin{},
}

// 路由
func (p *Admin) Route(app *fiber.App) {
	ag := app.Group("/api/admin")
	ag.Get("/login", (&controllers.Login{}).Show)
	ag.Post("/login", (&controllers.Login{}).Login)
	ag.Get("/logout", (&controllers.Login{}).Logout)
	ag.Get("/captcha", (&controllers.Captcha{}).Make)

	registerService(app)
}

// 注册服务
func registerService(app *fiber.App) {
	amg := app.Group("/api/admin", (&middleware.AdminMiddleware{}).Handle)

	// 仪表盘
	amg.Get("/dashboard/:dashboard", func(c *fiber.Ctx) error {
		var component interface{}
		providers := dashboardService
		for key, provider := range providers {
			if key == c.Params("dashboard") {
				component = provider.Render(c, provider, provider.DashboardComponentRender(c, provider))
			}
		}
		return c.JSON(component)
	})

	// 资源
	amg.Get("/:resource/index", func(c *fiber.Ctx) error {
		var component interface{}
		providers := resourceService
		for key, provider := range providers {
			if key == c.Params("resource") {
				component = provider.Render(c, provider, provider.IndexComponentRender())
			}
		}
		return c.JSON(component)
	})
}
