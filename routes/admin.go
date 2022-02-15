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

// 后台路由
func (p *Admin) Route(app *fiber.App) {

	login := &controllers.Login{}
	ag := app.Group("/api/admin")
	ag.Get("/login", login.Show)
	ag.Post("/login", login.Login)

	ag.Get("/captcha", (&controllers.Captcha{}).Make)

	registerRouteService(app)
}

// 注册路由服务
func registerRouteService(app *fiber.App) {
	adminMiddleware := &middleware.AdminMiddleware{}
	amg := app.Group("/api/admin", adminMiddleware.Handle)

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
