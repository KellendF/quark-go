package requests

import "github.com/gofiber/fiber/v2"

type ResourceEdit struct {
	Quark
}

// 表单数据
func (p *ResourceEdit) FillData(c *fiber.Ctx) map[string]interface{} {
	result := map[string]interface{}{}
	id := c.Query("id")
	if id == "" {
		return result
	}

	resourceInstance := p.Resource(c)
	model := p.NewModel(resourceInstance)
	model.Where("id = ?", id).First(&result)

	return result
}
