package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/http/controllers/admin/dashboard"
	"github.com/quarkcms/quark-go/internal/http/controllers/admin/login"
	"github.com/quarkcms/quark-go/internal/http/controllers/tool/captcha"
	"github.com/quarkcms/quark-go/internal/http/middleware"
)

func Admin(app *fiber.App) {

	// 后台路由
	ag := app.Group("/api/admin")
	ag.Get("/login", login.Show)
	ag.Post("/login", login.Login)
	ag.Get("/captcha", captcha.Make)

	// 后台认证路由
	amg := app.Group("/api/admin", middleware.Admin)
	amg.Get("/dashboard/index", dashboard.Index)
}
