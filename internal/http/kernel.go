package http

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/quarkcms/quark-go/internal/http/middleware"
	"github.com/quarkcms/quark-go/internal/providers"
	"github.com/quarkcms/quark-go/pkg/framework/config"
)

type Kernel struct{}

func (p *Kernel) Run(assets fs.FS) {

	if config.Get("app.debug").(string) == "false" {
		// 安装
		p.install()
	}

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

// 安装操作
func (p *Kernel) install() {

	// 创建软连接
	storagePath := filepath.Join("..", "storage", "app", "public")
	SymlinkPath := filepath.Join("public", "storage")

	err := os.Symlink(storagePath, SymlinkPath)
	if err != nil {
		fmt.Print(err)
	}
}
