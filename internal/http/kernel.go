package http

import (
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/quarkcms/quark-go/internal/http/middleware"
	"github.com/quarkcms/quark-go/internal/providers"
	"github.com/quarkcms/quark-go/pkg/framework/config"
)

type Kernel struct{}

func (p *Kernel) Run(assets fs.FS) {

	// 配置
	app := fiber.New(fiber.Config{
		AppName: config.Get("app.name").(string),
	})

	// 静态资源
	app.Static("/", "./public")

	// 中间件
	app.Use((&middleware.AppServiceInit{}).Handle)

	// 路由
	(&providers.Route{}).Register(app)

	// 资源打包到二进制文件
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(assets),
		Browse:     true,
		Index:      "index.html",
		MaxAge:     3600,
		PathPrefix: "assets",
	}))

	app.Listen(config.Get("app.host").(string))
}
