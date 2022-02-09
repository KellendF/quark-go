package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/http/middleware"
	"github.com/quarkcms/quark-go/internal/providers/route"
	"github.com/quarkcms/quark-go/pkg/framework/config"
)

type Kernel struct{}

func (p *Kernel) Run() {

	// 默认配置
	app := fiber.New(fiber.Config{
		AppName: config.Get("app.name"),
	})

	// 获取中间件
	appMiddleware := &middleware.App{}

	// 使用中间件
	app.Use(appMiddleware.Handle)

	// 注册路由
	route.Register(app)

	app.Listen(config.Get("app.host") + ":" + config.Get("app.port"))
}
