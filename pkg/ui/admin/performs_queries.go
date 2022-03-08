package admin

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// 创建列表查询
func (p *Resource) BuildIndexQuery(c *fiber.Ctx, resourceInstance interface{}, query *gorm.DB, search []interface{}, filters []interface{}, columnFilters interface{}, orderings interface{}) *gorm.DB {

	// 初始化查询
	query = p.initializeQuery(c, resourceInstance, query)

	// 执行列表查询，这里使用的是透传的实例
	query = resourceInstance.(interface {
		IndexQuery(*fiber.Ctx, *gorm.DB) *gorm.DB
	}).IndexQuery(c, query)

	// 执行搜索查询
	query = p.applySearch(c, query, search)

	// 执行过滤器查询
	query = p.applyFilters(query, filters)

	// 执行表格列上过滤器查询
	query = p.applyColumnFilters(query, filters)

	// 执行排序查询
	query = p.applyOrderings(query, filters)

	return query
}

// 创建详情页查询
func (p *Resource) BuildDetailQuery(c *fiber.Ctx, resourceInstance interface{}, query *gorm.DB, search interface{}, filters interface{}, columnFilters interface{}, orderings interface{}) interface{} {
	return "todo"
}

// 创建导出查询
func (p *Resource) BuildExportQuery(c *fiber.Ctx, resourceInstance interface{}, query *gorm.DB, search interface{}, filters interface{}, columnFilters interface{}, orderings interface{}) interface{} {
	return "todo"
}

// 初始化查询
func (p *Resource) initializeQuery(c *fiber.Ctx, resourceInstance interface{}, query *gorm.DB) *gorm.DB {

	return resourceInstance.(interface {
		Query(*fiber.Ctx, *gorm.DB) *gorm.DB
	}).Query(c, query)
}

// 执行搜索表单查询
func (p *Resource) applySearch(c *fiber.Ctx, query *gorm.DB, search []interface{}) *gorm.DB {

	for _, v := range search {

		// 获取字段
		column := reflect.
			ValueOf(v).
			Elem().
			FieldByName("Column").String()

		value := c.Query("search[" + column + "]")

		query = v.(interface {
			Apply(*fiber.Ctx, *gorm.DB, interface{}) *gorm.DB
		}).Apply(c, query, value)
	}

	return query
}

// 执行表格列上过滤器查询
func (p *Resource) applyColumnFilters(query *gorm.DB, filters interface{}) *gorm.DB {

	return query
}

// 执行过滤器查询
func (p *Resource) applyFilters(query *gorm.DB, filters []interface{}) *gorm.DB {

	return query
}

// 执行排序查询
func (p *Resource) applyOrderings(query *gorm.DB, orderings interface{}) *gorm.DB {

	return query
}

// 全局查询
func (p *Resource) Query(c *fiber.Ctx, query *gorm.DB) *gorm.DB {

	return query
}

// 列表查询
func (p *Resource) IndexQuery(c *fiber.Ctx, query *gorm.DB) *gorm.DB {

	return query
}

// 详情查询
func (p *Resource) DetailQuery(c *fiber.Ctx, query *gorm.DB) *gorm.DB {

	return query
}

// 导出查询
func (p *Resource) ExportQuery(c *fiber.Ctx, query *gorm.DB) *gorm.DB {

	return query
}
