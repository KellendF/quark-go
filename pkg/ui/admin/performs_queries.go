package admin

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// 创建列表查询
func (p *Resource) BuildIndexQuery(c *fiber.Ctx, resourceInstance interface{}, query *gorm.DB, search interface{}, filters interface{}, columnFilters interface{}, orderings interface{}) interface{} {
	var results []map[string]interface{}
	query.Find(&results)

	return results
}
