package admin

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/component/table"
)

// 列表页搜索表单
func (p *Resource) IndexSearches(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	searches := resourceInstance.(interface{ Searches(*fiber.Ctx) interface{} }).Searches(c)
	search := (&table.Search{}).Init()
	SearchItem := (&table.SearchItem{}).Init()

	withExport := reflect.
		ValueOf(resourceInstance).
		Elem().
		FieldByName("WithExport").Bool()

	if withExport {
		search = search.SetShowExportButton(true).SetExportApi("admin/" + c.Params("resource") + "/export")
	}

	for _, v := range searches.([]interface{}) {
		component := v.(interface{ GetComponent() string }).GetComponent() // 获取组件名称
		name := v.(interface{ GetName() string }).GetName()                // label 标签的文本
		column := v.(interface{ GetColumn() string }).GetColumn()          // 字段名，支持数组
		operator := v.(interface{ GetOperator() string }).GetOperator()    // 获取操作符
		api := v.(interface{ GetApi() string }).GetApi()                   // 获取接口
		options := v.(interface {
			Options(c *fiber.Ctx) interface{}
		}).Options(c) // 获取属性

		// 搜索栏表单项
		SearchItem = SearchItem.SetName(column).SetLabel(name).SetOperator(operator).SetApi(api)

		switch component {
		case "input":
			SearchItem = SearchItem.Input(options)
		case "select":
			SearchItem = SearchItem.Select(options.([]interface{}))
		case "multipleSelect":
			SearchItem = SearchItem.MultipleSelect(options.([]interface{}))
		case "datetime":
			SearchItem = SearchItem.Datetime(options)
		case "date":
			SearchItem = SearchItem.Date(options)
		case "cascader":
			SearchItem = SearchItem.Cascader(options.([]interface{}))
		}

		search = search.SetItems(SearchItem)
	}

	return search
}
