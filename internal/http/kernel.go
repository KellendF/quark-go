package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/http/middleware"
	"github.com/quarkcms/quark-go/internal/providers"
	"github.com/quarkcms/quark-go/pkg/framework/config"
)

type Kernel struct{}

func (p *Kernel) Run() {
	appMiddleware := &middleware.App{}
	route := &providers.Route{}

	// 默认配置
	app := fiber.New(fiber.Config{
		AppName: config.Get("app.name"),
	})

	// 使用中间件
	app.Use(appMiddleware.Handle)

	// 注册路由
	route.Register(app)

	app.Listen(config.Get("app.host") + ":" + config.Get("app.port"))
}
