package requests

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ResourceIndex struct {
	Quark
}

// 列表查询
func (p *ResourceIndex) IndexQuery(c *fiber.Ctx) []map[string]interface{} {
	var results []map[string]interface{}

	resourceInstance := p.Resource(c)
	model := p.NewModel(resourceInstance)

	query := resourceInstance.(interface {
		BuildIndexQuery(c *fiber.Ctx, resourceInstance interface{}, query *gorm.DB, search interface{}, filters interface{}, columnFilters interface{}, orderings interface{}) *gorm.DB
	}).BuildIndexQuery(c, resourceInstance, model, "", "", "", "")

	query.Find(&results)

	return results
}
