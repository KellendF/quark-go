package requests

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ResourceAction struct {
	Quark
}

// 执行行为
func (p *ResourceAction) HandleAction(c *fiber.Ctx) error {
	var result error
	resourceInstance := p.Resource(c)
	model := p.NewModel(resourceInstance)

	id := c.Query("id")

	if id != "" {
		if strings.Contains(id, ",") {
			model.Where("id IN ?", strings.Split(",", id))
		} else {
			model.Where("id = ?", id)
		}
	}

	actions := resourceInstance.(interface {
		Actions(c *fiber.Ctx) interface{}
	}).Actions(c)

	for _, v := range actions.([]interface{}) {

		// uri唯一标识
		uriKey := v.(interface {
			GetUriKey(interface{}) string
		}).GetUriKey(v)

		actionType := v.(interface{ GetActionType() string }).GetActionType()

		if actionType == "dropdown" {
			// todo
		} else {
			if c.Params("uriKey") == uriKey {
				result = v.(interface {
					Handle(*fiber.Ctx, *gorm.DB) error
				}).Handle(c, model)
			}
		}
	}

	return result
}
