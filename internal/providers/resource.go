package providers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/http/controllers/admin/dashboards"
	"github.com/quarkcms/quark-go/internal/http/controllers/admin/resources"
	"github.com/quarkcms/quark-go/internal/http/middleware"
	"github.com/quarkcms/quark-go/pkg/ui/admin/dashboard"
	"github.com/quarkcms/quark-go/pkg/ui/admin/resource"
)

// 结构体
type Resource struct{}

// 仪表盘服务
var dashboardService = map[string]dashboard.DashboardInterface{
	"index": &dashboards.Index{},
}

// 资源服务
var resourceService = map[string]resource.ResourceInterface{
	"admin": &resources.Admin{},
}

// 仪表盘路由
func (p *Resource) DashboardRoute(app *fiber.App) {
	adminMiddleware := &middleware.AdminMiddleware{}
	amg := app.Group("/api/admin", adminMiddleware.Handle)
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
}

// 资源路由
func (p *Resource) ResourceRoute(app *fiber.App) {
	adminMiddleware := &middleware.AdminMiddleware{}
	amg := app.Group("/api/admin", adminMiddleware.Handle)
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
