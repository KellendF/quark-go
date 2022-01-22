package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/http/middleware"
	"github.com/quarkcms/quark-go/routes"
)

// 注册服务
func Register(app *fiber.App) {

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
	amg := app.Group("/api/admin", middleware.Admin)
	amg.Get("/dashboard/:dashboard", func(c *fiber.Ctx) error {
		var component interface{}
		providers := routes.Dashboard()

		for key, provider := range providers {
			if key == c.Params("dashboard") {
				component = provider.Render(c, provider, provider.DashboardComponentRender())
			}
		}

		return c.JSON(component)
	})
}

// 加载资源路由
func loadResourceRoute(app *fiber.App) {
	amg := app.Group("/api/admin", middleware.Admin)
	amg.Get("/:resource/index", func(c *fiber.Ctx) error {
		var component interface{}
		providers := routes.Resource()

		for key, provider := range providers {
			if key == c.Params("resource") {
				component = provider.Render(c, provider, provider.IndexComponentRender())
			}
		}

		return c.JSON(component)
	})
}
