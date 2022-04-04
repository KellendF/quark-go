package controllers

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/admin/http/requests"
)

type ResourceExport struct{}

// 导出数据
func (p *ResourceExport) Handle(c *fiber.Ctx) error {

	resourceExport := &requests.ResourceExport{}

	// 资源实例
	resourceInstance := resourceExport.Resource(c)

	if resourceInstance == nil {
		return c.SendStatus(404)
	}

	data := resourceExport.IndexQuery(c)

	// 获取列表字段
	fields := resourceInstance.(interface {
		ExportFields(c *fiber.Ctx, resourceInstance interface{}) interface{}
	}).ExportFields(c, resourceInstance)

	exportData := []map[string]interface{}{}
	exportTitles := []string{}

	rowData := map[string]interface{}{}

	for _, fieldValue := range fields.([]interface{}) {
		Label := reflect.
			ValueOf(fieldValue).
			Elem().
			FieldByName("Label").
			String()

		exportTitles = append(exportTitles, Label)
	}

	for _, dataValue := range data.([]interface{}) {

		for _, fieldValue := range fields.([]interface{}) {

			name := reflect.
				ValueOf(fieldValue).
				Elem().
				FieldByName("Name").String()

			component := reflect.
				ValueOf(fieldValue).
				Elem().
				FieldByName("Component").String()

			switch component {
			case "inputNumberField":
				rowData[name] = dataValue.(map[string]interface{})[name]

			case "textField":
				rowData[name] = dataValue.(map[string]interface{})[name]

			case "selectField":
				options := reflect.
					ValueOf(fieldValue).
					Elem().
					FieldByName("Options").Interface()

				rowData[name] = p.getOptionValue(options, dataValue.(map[string]interface{})[name])

			case "cascaderField":
				options := reflect.
					ValueOf(fieldValue).
					Elem().
					FieldByName("Options").Interface()

				rowData[name] = p.getOptionValue(options, dataValue.(map[string]interface{})[name])

			case "checkboxField":
				options := reflect.
					ValueOf(fieldValue).
					Elem().
					FieldByName("Options").Interface()

				rowData[name] = p.getOptionValue(options, dataValue.(map[string]interface{})[name])

			case "radioField":
				options := reflect.
					ValueOf(fieldValue).
					Elem().
					FieldByName("Options").Interface()

				rowData[name] = p.getOptionValue(options, dataValue.(map[string]interface{})[name])

			case "switchField":
				options := reflect.
					ValueOf(fieldValue).
					Elem().
					FieldByName("Options").Interface()

				rowData[name] = p.getSwitchValue(options, dataValue.(map[string]interface{})[name].(int))

			default:
				rowData[name] = dataValue.(map[string]interface{})[name]
			}
		}
		exportData = append(exportData, rowData)
	}

	return c.JSON(map[string]interface{}{
		"a": exportTitles, "b": exportData,
	})
}

// 获取属性值
func (p *ResourceExport) getOptionValue(options interface{}, value interface{}) string {
	result := ""
	arr := []interface{}{}

	if value, ok := value.(string); ok {
		if strings.Contains(value, "[") || strings.Contains(value, "{") {
			json.Unmarshal([]byte(value), arr)
		}
	}

	if len(arr) > 0 {
		for _, option := range options.([]interface{}) {
			for _, v := range arr {
				if v == option.(map[string]interface{})["value"] {
					result = result + option.(map[string]interface{})["label"].(string)
				}
			}
		}
	} else {

		for _, option := range options.([]interface{}) {
			if value.(string) == option.(map[string]interface{})["value"] {
				result = option.(map[string]interface{})["label"].(string)
			}
		}
	}

	return result
}

// 获取开关组件值
func (p *ResourceExport) getSwitchValue(options interface{}, value int) string {
	if value == 1 {
		return options.(map[string]interface{})["on"].(string)
	} else {
		return options.(map[string]interface{})["off"].(string)
	}
}
