package admin

import (
	"github.com/gofiber/fiber/v2"
)

// 列表字段
func (p *Resource) IndexFields(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	fields := p.getFields(c, resourceInstance)
	var items []interface{}

	for _, v := range fields.([]interface{}) {

		isShownOnIndex := v.(interface {
			IsShownOnIndex() bool
		}).IsShownOnIndex()

		if isShownOnIndex {
			items = append(items, v)
		}
	}

	return items
}

// 表格列
func (p *Resource) IndexColumns(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	fields := p.IndexFields(c, resourceInstance)
	var columns []interface{}

	for _, v := range fields.([]interface{}) {

		isShownOnIndex := v.(interface {
			IsShownOnIndex() bool
		}).IsShownOnIndex()

		if isShownOnIndex {
			column := v.(interface {
				TransformToColumn() interface{}
			}).TransformToColumn()

			columns = append(columns, column)
		}
	}

	return columns
}

// 获取字段
func (p *Resource) getFields(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	return resourceInstance.(interface {
		Fields(c *fiber.Ctx) interface{}
	}).Fields(c)
}
