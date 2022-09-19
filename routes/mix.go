package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/mix/controllers"
	"github.com/quarkcms/quark-go/internal/mix/middleware"
)

type Mix struct{}

// mix路由
func (p *Mix) Route(app *fiber.App) {

	ag := app.Group("/api/mix")
	ag.Get("/index/index", (&controllers.Index{}).Index)

	amg := app.Group("/api/mix", (&middleware.ApiMiddleware{}).Handle)
	amg.Get("/user/index", (&controllers.User{}).Index)
}
