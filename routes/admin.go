package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/admin/http/controllers"
	"github.com/quarkcms/quark-go/pkg/ui/admin/http/middleware"
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
	amg.Get("/dashboard/:dashboard", (&controllers.Dashboard{}).Handle)          // 仪表盘
	amg.Get("/:resource/index", (&controllers.ResourceIndex{}).Handle)           // 资源
	amg.Get("/:resource/editable", (&controllers.ResourceEditable{}).Handle)     // 表格行内编辑
	amg.Get("/:resource/action/:uriKey", (&controllers.ResourceAction{}).Handle) // 执行行为
}
