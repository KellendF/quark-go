package requests

import (
	"encoding/json"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ResourceStore struct {
	Quark
}

// 执行行为
func (p *ResourceStore) HandleStore(c *fiber.Ctx) interface{} {
	resourceInstance := p.Resource(c)
	model := p.NewModel(resourceInstance)

	// 获取字段
	fields := resourceInstance.(interface {
		CreationFields(c *fiber.Ctx, resourceInstance interface{}) interface{}
	}).CreationFields(c, resourceInstance)

	data := map[string]interface{}{}
	json.Unmarshal(c.Body(), &data)

	getData := resourceInstance.(interface {
		BeforeSaving(c *fiber.Ctx, data map[string]interface{}) interface{}
	}).BeforeSaving(c, data)

	if value, ok := getData.(error); ok {
		return value
	}

	if value, ok := getData.(map[string]interface{}); ok {
		data = value
	}

	for _, v := range fields.([]interface{}) {

		name := reflect.
			ValueOf(v).
			Elem().
			FieldByName("Name").String()

		if name == "" {
			delete(data, name)
		}

		arrayValue, ok := data[name].([]interface{})

		if ok {
			data[name], _ = json.Marshal(arrayValue)
		}
	}

	validator := resourceInstance.(interface {
		ValidatorForCreation(c *fiber.Ctx, resourceInstance interface{}, data map[string]interface{}) error
	}).ValidatorForCreation(c, resourceInstance, data)

	if validator != nil {
		return validator
	}

	// 获取对象
	getModel := model.Create(data)

	return resourceInstance.(interface {
		AfterSaved(c *fiber.Ctx, model *gorm.DB) interface{}
	}).AfterSaved(c, getModel)
}
