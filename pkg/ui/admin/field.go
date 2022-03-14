package admin

import (
	"github.com/quarkcms/quark-go/pkg/ui/component/form/fields"
)

type Field struct{}

// 输入框组件
func (p *Field) Text(params ...interface{}) *fields.Text {
	field := &fields.Text{}

	if len(params) >= 2 {
		field.Init().SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		field.Init().SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 单选组件
func (p *Field) Radio(params ...string) *fields.Radio {
	field := &fields.Radio{}

	if len(params) == 2 {
		field.Init().SetName(params[0]).SetLabel(params[1])
	} else {
		field.Init().SetName(params[0]).SetLabel(params[0])
	}

	return field
}

// 日期时间组件
func (p *Field) Datetime(params ...string) *fields.Datetime {
	field := &fields.Datetime{}

	if len(params) == 2 {
		field.Init().SetName(params[0]).SetLabel(params[1])
	} else {
		field.Init().SetName(params[0]).SetLabel(params[0])
	}

	return field
}

// 开关组件
func (p *Field) Switch(params ...string) *fields.Switch {
	field := &fields.Switch{}

	if len(params) == 2 {
		field.Init().SetName(params[0]).SetLabel(params[1])
	} else {
		field.Init().SetName(params[0]).SetLabel(params[0])
	}

	return field
}
