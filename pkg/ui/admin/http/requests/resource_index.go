package requests

import (
	"reflect"
	"strconv"

	"github.com/derekstavis/go-qs"
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
		BuildIndexQuery(c *fiber.Ctx, resourceInstance interface{}, query *gorm.DB, search []interface{}, filters []interface{}, columnFilters map[string]interface{}, orderings map[string]interface{}) *gorm.DB
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

	// 获取列表字段
	indexFields := resourceInstance.(interface {
		IndexFields(c *fiber.Ctx, resourceInstance interface{}) interface{}
	}).IndexFields(c, resourceInstance)

	// 解析字段回调函数
	for k, v := range lists {

		// 给实例的Field属性赋值
		resourceInstance.(interface {
			SetField(fieldData map[string]interface{}) interface{}
		}).SetField(v)

		fields := make(map[string]interface{})
		for _, field := range indexFields.([]interface{}) {

			// 字段名
			name := reflect.
				ValueOf(field).
				Elem().
				FieldByName("Name").String()

			// 获取实例的回调函数
			callback := field.(interface{ GetCallback() interface{} }).GetCallback()

			if callback != nil {
				getCallback := callback.(func() interface{})
				fields[name] = getCallback()
			} else {
				if v[name] != nil {
					fields[name] = v[name]
				}
			}
		}

		lists[k] = fields
	}

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
func (p *ResourceIndex) columnFilters(c *fiber.Ctx) map[string]interface{} {
	data, error := qs.Unmarshal(c.OriginalURL())

	if error != nil {
		return map[string]interface{}{}
	}

	result, ok := data["filter"].(map[string]interface{})

	if ok == false {
		return map[string]interface{}{}
	}

	return result
}

/**
 * Get the orderings for the request.
 *
 * @return array
 */
func (p *ResourceIndex) orderings(c *fiber.Ctx) map[string]interface{} {
	data, error := qs.Unmarshal(c.OriginalURL())

	if error != nil {
		return map[string]interface{}{}
	}

	result, ok := data["sorter"].(map[string]interface{})

	if ok == false {
		return map[string]interface{}{}
	}

	return result
}
