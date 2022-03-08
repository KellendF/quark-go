package requests

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ResourceEditable struct {
	Quark
}

// 执行行为
func (p *ResourceEditable) HandleEditable(c *fiber.Ctx) interface{} {
	resourceInstance := p.Resource(c)
	model := p.NewModel(resourceInstance)

	// todo
	fmt.Println(c.GetReqHeaders())

	return model
}
