package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/http/middleware"
	"github.com/quarkcms/quark-go/pkg/config"
	"github.com/quarkcms/quark-go/routes"
)

func Run() {

	// 默认配置
	app := fiber.New(fiber.Config{
		AppName: config.Get("app.name"),
	})

	// 全局中间件
	app.Use(middleware.App)

	// 注册Admin路由
	routes.Admin(app)

	// 注册Web路由
	routes.Web(app)

	app.Listen(config.Get("app.host") + ":" + config.Get("app.port"))
}
