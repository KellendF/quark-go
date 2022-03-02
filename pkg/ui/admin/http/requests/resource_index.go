package requests

import (
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/admin"
	"gorm.io/gorm"
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

				// 初始化实例
				resourceInstance := provider.(interface{ Init() interface{} }).Init()

				// 获取绑定的模型
				model := resourceInstance.(interface{ NewModel(interface{}) *gorm.DB }).NewModel(resourceInstance)

				// 查询列表数据
				data := resourceInstance.(interface {
					BuildIndexQuery(c *fiber.Ctx, resourceInstance interface{}, query *gorm.DB, search interface{}, filters interface{}, columnFilters interface{}, orderings interface{}) interface{}
				}).
					BuildIndexQuery(c, resourceInstance, model, "", "", "", "")

				// 断言IndexComponentRender方法
				indexComponent := provider.(interface {
					IndexComponentRender(*fiber.Ctx, interface{}, interface{}) interface{}
				}).IndexComponentRender(c, resourceInstance, data)

				// 断言Render方法
				component = provider.(interface {
					Render(*fiber.Ctx, interface{}, interface{}) interface{}
				}).Render(c, resourceInstance, indexComponent)
			}
		}
	}

	return c.JSON(component)
}
