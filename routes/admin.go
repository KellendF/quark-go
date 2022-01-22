package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/http/controllers/admin/login"
	"github.com/quarkcms/quark-go/internal/http/controllers/tool/captcha"
)

// 后台路由
func Admin(app *fiber.App) {
	ag := app.Group("/api/admin")
	ag.Get("/login", login.Show)
	ag.Post("/login", login.Login)
	ag.Get("/captcha", captcha.Make)
}
