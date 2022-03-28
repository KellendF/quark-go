package requests

import (
	"encoding/json"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/framework/msg"
	"gorm.io/gorm"
)

type ResourceUpdate struct {
	Quark
}

// 执行行为
func (p *ResourceUpdate) HandleUpdate(c *fiber.Ctx) interface{} {
	resourceInstance := p.Resource(c)
	model := p.NewModel(resourceInstance)

	// 获取字段
	fields := resourceInstance.(interface {
		UpdateFields(c *fiber.Ctx, resourceInstance interface{}) interface{}
	}).UpdateFields(c, resourceInstance)

	data := map[string]interface{}{}
	json.Unmarshal(c.Body(), &data)

	if data["id"] == "" {
		return msg.Error("参数错误！", "")
	}

	getData := resourceInstance.(interface {
		BeforeSaving(c *fiber.Ctx, data map[string]interface{}) interface{}
	}).BeforeSaving(c, data)

	if value, ok := getData.(error); ok {
		return value
	}

	if value, ok := getData.(map[string]interface{}); ok {
		data = value
	}

	validator := resourceInstance.(interface {
		ValidatorForUpdate(c *fiber.Ctx, resourceInstance interface{}, data map[string]interface{}) error
	}).ValidatorForUpdate(c, resourceInstance, data)

	if validator != nil {
		return validator
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

	// 获取对象
	getModel := model.Where("id = ?", data["id"]).Updates(data)

	return resourceInstance.(interface {
		AfterSaved(c *fiber.Ctx, model *gorm.DB) interface{}
	}).AfterSaved(c, getModel)
}
