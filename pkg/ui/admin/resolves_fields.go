package admin

import (
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/component/table"
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
			columns = append(columns, p.fieldToColumn(c, v))
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

// 表单项转换为表格列
func (p *Resource) fieldToColumn(c *fiber.Ctx, field interface{}) interface{} {

	// 字段
	name := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Name").String()

	// 文字
	label := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Label").String()

	// 组件类型
	component := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Component").String()

	// 是否可编辑
	editable := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Editable").Bool()

	column := (&table.Column{}).
		Init().
		SetTitle(label).
		SetAttribute(name)

	switch component {
	case "textField":
		column = column.SetValueType("text")
	case "selectField":
		column = column.SetValueType("select")
	case "radioField":
		column = column.SetValueType("radio")
	case "switchField":
		column = column.SetValueType("select")
	case "imageField":
		column = column.SetValueType("image")
	default:
		column = column.SetValueType(component)
	}

	if editable {
		// 可编辑，设置编辑
		options := reflect.
			ValueOf(field).
			Elem().
			FieldByName("Options").Interface()

		// 可编辑api地址
		editableApi := strings.Replace(strings.Replace(c.Path(), "/api/", "", -1), "/index", "/editable", -1)

		// 设置编辑项
		column = column.SetEditable(component, options, editableApi)
	}

	return column
}
