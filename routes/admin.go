package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/admin/http/controllers"
	"github.com/quarkcms/quark-go/pkg/ui/admin/http/middleware"
	"github.com/quarkcms/quark-go/pkg/ui/admin/http/requests"
)

type Admin struct{}

// 路由
func (p *Admin) Route(app *fiber.App) {
	ag := app.Group("/api/admin")
	ag.Get("/login", (&controllers.Login{}).Show)
	ag.Post("/login", (&controllers.Login{}).Login)
	ag.Get("/logout", (&controllers.Login{}).Logout)
	ag.Get("/captcha", (&controllers.Captcha{}).Make)

	amg := app.Group("/api/admin", (&middleware.AdminMiddleware{}).Handle)
	amg.Get("/dashboard/:dashboard", (&requests.Dashboard{}).Resource) // 仪表盘
	amg.Get("/:resource/index", (&requests.ResourceIndex{}).Resource)  // 资源
}
