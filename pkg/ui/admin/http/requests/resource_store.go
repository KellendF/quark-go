package requests

import (
	"encoding/json"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type ResourceStore struct {
	Quark
}

// 执行行为
func (p *ResourceStore) HandleStore(c *fiber.Ctx) error {
	resourceInstance := p.Resource(c)
	model := p.NewModel(resourceInstance)
	data := map[string]interface{}{}

	// 获取字段
	fields := resourceInstance.(interface {
		CreationFields(c *fiber.Ctx, resourceInstance interface{}) interface{}
	}).CreationFields(c, resourceInstance)

	json.Unmarshal(c.Body(), &data)

	for _, v := range fields.([]interface{}) {

		name := reflect.
			ValueOf(v).
			Elem().
			FieldByName("Name").String()

		if name == "" {
			delete(data, name)
		}
	}

	return model.Create(data).Error
}
