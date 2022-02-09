package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/http/controllers/admin/resources"
	"github.com/quarkcms/quark-go/internal/http/controllers/tool"
)

// 后台路由
func Admin(app *fiber.App) {

	login := &resources.Login{}
	ag := app.Group("/api/admin")
	ag.Get("/login", login.Show)
	ag.Post("/login", login.Login)

	captcha := &tool.Captcha{}
	ag.Get("/captcha", captcha.Make)
}
