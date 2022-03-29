package admin

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/ui/component/form/fields"
)

type Field struct{}

// ID组件
func (p *Field) ID(params ...interface{}) *fields.ID {
	field := (&fields.ID{}).Init()

	if len(params) >= 2 {

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// Hidden组件
func (p *Field) Hidden(params ...interface{}) *fields.Hidden {
	field := (&fields.Hidden{}).Init()

	if len(params) >= 2 {

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 输入框组件
func (p *Field) Text(params ...interface{}) *fields.Text {
	field := (&fields.Text{}).Init()

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.SetPlaceholder("请输入" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.SetPlaceholder("请输入" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 密码框组件
func (p *Field) Password(params ...interface{}) *fields.Password {
	field := (&fields.Password{}).Init()

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.SetPlaceholder("请输入" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.SetPlaceholder("请输入" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
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

// 多选组件
func (p *Field) Checkbox(params ...string) *fields.Checkbox {
	field := &fields.Checkbox{}

	if len(params) == 2 {
		field.Init().SetName(params[0]).SetLabel(params[1])
	} else {
		field.Init().SetName(params[0]).SetLabel(params[0])
	}

	return field
}

// 日期时间组件
func (p *Field) Datetime(params ...interface{}) *fields.Datetime {
	field := &fields.Datetime{}

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.Init().SetPlaceholder("请选择")
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
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

// 输入框组件
func (p *Field) Tree(params ...interface{}) *fields.Tree {
	field := (&fields.Tree{}).Init()

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.SetPlaceholder("请选择" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.SetPlaceholder("请选择" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 图标组件
func (p *Field) Icon(params ...interface{}) *fields.Icon {
	field := (&fields.Icon{}).Init()

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.SetPlaceholder("请选择" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.SetPlaceholder("请选择" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 图标组件
func (p *Field) Select(params ...interface{}) *fields.Select {
	field := (&fields.Select{}).Init()

	placeholder := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Placeholder").String()

	if len(params) >= 2 {

		if placeholder == "" {
			field.SetPlaceholder("请选择" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {
		if placeholder == "" {
			field.SetPlaceholder("请选择" + params[1].(string))
		}

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}

// 图片组件
func (p *Field) Image(params ...interface{}) *fields.Image {
	field := (&fields.Image{}).Init()

	if len(params) >= 2 {

		field.SetName(params[0].(string)).SetLabel(params[1].(string))
		if len(params) == 3 {

			// 判断是否为闭包函数
			closure, ok := params[2].(func() interface{})
			if ok {
				field.SetCallback(closure)
			}
		}
	} else {

		field.SetName(params[0].(string)).SetLabel(params[0].(string))
	}

	return field
}
