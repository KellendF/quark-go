package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/http/middleware"
	"github.com/quarkcms/quark-go/internal/providers"
	"github.com/quarkcms/quark-go/pkg/framework/config"
)

type Kernel struct{}

func (p *Kernel) Run() {

	// 配置
	app := fiber.New(fiber.Config{
		AppName: config.Get("app.name").(string),
	})

	// 中间件
	app.Use((&middleware.AppServiceInit{}).Handle)

	// 路由
	(&providers.Route{}).Register(app)

	app.Listen(config.Get("app.host").(string))
}
