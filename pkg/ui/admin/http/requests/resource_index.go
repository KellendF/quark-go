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

	// 搜索项
	searches := resourceInstance.(interface {
		Searches(c *fiber.Ctx) []interface{}
	}).Searches(c)

	// 过滤项，预留
	filters := resourceInstance.(interface {
		Filters(c *fiber.Ctx) []interface{}
	}).Filters(c)

	query := resourceInstance.(interface {
		BuildIndexQuery(c *fiber.Ctx, resourceInstance interface{}, query *gorm.DB, search []interface{}, filters []interface{}, columnFilters interface{}, orderings interface{}) *gorm.DB
	}).BuildIndexQuery(c, resourceInstance, model, searches, filters, p.columnFilters(c), p.orderings(c))

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
	page := c.Query("page", "1")
	pageSize := c.Query("pageSize")

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

/**
 * Get the column filters for the request.
 *
 * @return array
 */
func (p *ResourceIndex) columnFilters(c *fiber.Ctx) interface{} {
	filter := c.Params("filter")

	if filter != "" {
		return filter
	} else {
		return nil
	}
}

/**
 * Get the orderings for the request.
 *
 * @return array
 */
func (p *ResourceIndex) orderings(c *fiber.Ctx) interface{} {
	sorter := c.Params("sorter")

	if sorter != "" {
		return sorter
	} else {
		return nil
	}
}
