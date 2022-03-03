package requests

import (
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ResourceIndex struct {
	Quark
}

// 列表查询
func (p *ResourceIndex) IndexQuery(c *fiber.Ctx) interface{} {
	var lists []map[string]interface{}
	resourceInstance := p.Resource(c)
	model := p.NewModel(resourceInstance)

	query := resourceInstance.(interface {
		BuildIndexQuery(c *fiber.Ctx, resourceInstance interface{}, query *gorm.DB, search interface{}, filters interface{}, columnFilters interface{}, orderings interface{}) *gorm.DB
	}).BuildIndexQuery(c, resourceInstance, model, "", "", "", "")

	// 获取分页
	perPage := reflect.
		ValueOf(resourceInstance).
		Elem().
		FieldByName("PerPage").Interface()

	// 不分页，直接返回lists
	if reflect.TypeOf(perPage).String() != "int" {
		query.Find(&lists)
		return lists
	}

	var total int64
	page := c.Params("page", "1")
	pageSize := c.Params("pageSize")

	if pageSize != "" {
		perPage, _ = strconv.Atoi(pageSize)
	}

	getPage, _ := strconv.Atoi(page)
	query.Limit(perPage.(int)).Offset((getPage - 1) * perPage.(int)).Find(&lists).Count(&total)

	return map[string]interface{}{
		"currentPage": getPage,
		"perPage":     perPage,
		"total":       total,
		"items":       lists,
	}
}
