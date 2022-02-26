package routes

import (
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/admin/dashboards"
	"github.com/quarkcms/quark-go/internal/admin/resources"
	"github.com/quarkcms/quark-go/pkg/ui/admin/http/controllers"
	"github.com/quarkcms/quark-go/pkg/ui/admin/http/middleware"
)

type Admin struct{}

// providers
var providers = []interface{}{
	&dashboards.Index{},
	&resources.Admin{},
	&resources.MerchantCategory{},
}

// 路由
func (p *Admin) Route(app *fiber.App) {
	ag := app.Group("/api/admin")
	ag.Get("/login", (&controllers.Login{}).Show)
	ag.Post("/login", (&controllers.Login{}).Login)
	ag.Get("/logout", (&controllers.Login{}).Logout)
	ag.Get("/captcha", (&controllers.Captcha{}).Make)

	registerService(app)
}

// 注册服务
func registerService(app *fiber.App) {
	amg := app.Group("/api/admin", (&middleware.AdminMiddleware{}).Handle)

	// 仪表盘
	amg.Get("/dashboard/:dashboard", func(c *fiber.Ctx) error {
		var component interface{}
		for _, provider := range providers {

			providerName := reflect.TypeOf(provider).String()

			// 包含字符串
			if find := strings.Contains(providerName, "*dashboards."); find {
				structName := strings.Replace(providerName, "*dashboards.", "", -1)

				if strings.ToLower(structName) == strings.ToLower(c.Params("dashboard")) {

					// 断言DashboardComponentRender方法
					dashboardComponent := provider.(interface {
						DashboardComponentRender(*fiber.Ctx, interface{}) interface{}
					}).DashboardComponentRender(c, provider)

					// 断言Render方法
					component = provider.(interface {
						Render(*fiber.Ctx, interface{}, interface{}) interface{}
					}).Render(c, provider, dashboardComponent)
				}
			}
		}

		return c.JSON(component)
	})

	// 资源
	amg.Get("/:resource/index", func(c *fiber.Ctx) error {
		var component interface{}
		for _, provider := range providers {
			providerName := reflect.TypeOf(provider).String()

			// 包含字符串
			if find := strings.Contains(providerName, "*resources."); find {
				structName := strings.Replace(providerName, "*resources.", "", -1)

				if strings.ToLower(structName) == strings.ToLower(c.Params("resource")) {

					// 断言IndexComponentRender方法
					indexComponent := provider.(interface {
						IndexComponentRender(*fiber.Ctx, interface{}) interface{}
					}).IndexComponentRender(c, provider)

					// 断言Render方法
					component = provider.(interface {
						Render(*fiber.Ctx, interface{}, interface{}) interface{}
					}).Render(c, provider, indexComponent)
				}
			}
		}

		return c.JSON(component)
	})
}
