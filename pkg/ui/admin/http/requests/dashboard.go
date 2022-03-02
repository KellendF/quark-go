package requests

import (
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/admin"
)

type Dashboard struct{}

func (p *Dashboard) Resource(c *fiber.Ctx) error {
	var component interface{}
	for _, provider := range admin.Providers {

		providerName := reflect.TypeOf(provider).String()

		// 包含字符串
		if find := strings.Contains(providerName, "*dashboards."); find {
			structName := strings.Replace(providerName, "*dashboards.", "", -1)

			if strings.ToLower(structName) == strings.ToLower(c.Params("dashboard")) {

				// 初始化实例
				dashboardInstance := provider.(interface{ Init() interface{} }).Init()

				// 断言DashboardComponentRender方法
				dashboardComponent := provider.(interface {
					DashboardComponentRender(*fiber.Ctx, interface{}) interface{}
				}).DashboardComponentRender(c, dashboardInstance)

				// 断言Render方法
				component = provider.(interface {
					Render(*fiber.Ctx, interface{}, interface{}) interface{}
				}).Render(c, dashboardInstance, dashboardComponent)
			}
		}
	}

	return c.JSON(component)
}
