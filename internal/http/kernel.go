package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/http/middleware"
	"github.com/quarkcms/quark-go/internal/http/providers/route"
	"github.com/quarkcms/quark-go/pkg/config"
)

func Run() {

	// 默认配置
	app := fiber.New(fiber.Config{
		AppName: config.Get("app.name"),
	})

	// 中间件
	app.Use(middleware.App)

	// 注册路由
	route.Register(app)

	app.Listen(config.Get("app.host") + ":" + config.Get("app.port"))
}
