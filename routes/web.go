package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/http/controllers/tool/captcha"
)

// web路由
func Web(app *fiber.App) {
	tg := app.Group("/tool")
	tg.Get("/captcha/getId", captcha.GetID)
	tg.Get("/captcha/id/:id?", captcha.MakeByID)
	tg.Get("/captcha/session", captcha.Make)
}
