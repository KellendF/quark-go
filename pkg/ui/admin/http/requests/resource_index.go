package requests

import (
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/admin"
)

type ResourceIndex struct{}

func (p *ResourceIndex) Resource(c *fiber.Ctx) error {
	var component interface{}
	for _, provider := range admin.Providers {
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
}
