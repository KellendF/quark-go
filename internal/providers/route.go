package providers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/http/middleware"
	"github.com/quarkcms/quark-go/routes"
	"github.com/quarkcms/quark-go/routes/resource"
)

// 结构体
type Route struct{}

// 注册服务
func (p *Route) Register(app *fiber.App) {

	// 注册Admin路由
	routes.Admin(app)

	// 注册Web路由
	routes.Web(app)

	// 加载仪表盘路由
	loadDashboardRoute(app)

	// 加载资源路由
	loadResourceRoute(app)
}

// 加载仪表盘路由
func loadDashboardRoute(app *fiber.App) {
	adminMiddleware := &middleware.Admin{}
	amg := app.Group("/api/admin", adminMiddleware.Handle)
	amg.Get("/dashboard/:dashboard", func(c *fiber.Ctx) error {
		var component interface{}
		providers := resource.Dashboard()

		for key, provider := range providers {
			if key == c.Params("dashboard") {
				component = provider.Render(c, provider, provider.DashboardComponentRender(c, provider))
			}
		}

		return c.JSON(component)
	})
}

// 加载资源路由
func loadResourceRoute(app *fiber.App) {
	adminMiddleware := &middleware.Admin{}
	amg := app.Group("/api/admin", adminMiddleware.Handle)
	amg.Get("/:resource/index", func(c *fiber.Ctx) error {
		var component interface{}
		providers := resource.Resource()

		for key, provider := range providers {
			if key == c.Params("resource") {
				component = provider.Render(c, provider, provider.IndexComponentRender())
			}
		}

		return c.JSON(component)
	})
}
