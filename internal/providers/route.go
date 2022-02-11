package providers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/routes"
)

// 结构体
type Route struct{}

// 注册服务
func (p *Route) Register(app *fiber.App) {

	// 注册Admin路由
	routes.Admin(app)

	// 注册Web路由
	routes.Web(app)

	resource := &Resource{}

	// 加载仪表盘路由
	resource.DashboardRoute(app)

	// 加载资源路由
	resource.ResourceRoute(app)
}
