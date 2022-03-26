package fields

import "github.com/quarkcms/quark-go/pkg/ui/component"

type Select struct {
	Item
}

// 初始化
func (p *Select) Init() *Select {
	p.Component = "selectField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.SetWidth(200)

	return p
}

// 单向联动
func (p *Select) SetLoad(field string, api string) *Select {
	p.Load = map[string]string{
		"field": field,
		"api":   api,
	}

	return p
}

// 设置单选属性
func (p *Select) SetOptions(options map[interface{}]interface{}) *Select {
	var data []map[string]interface{}
	for k, v := range options {
		option := map[string]interface{}{
			"label": v,
			"value": k,
		}

		data = append(data, option)
	}

	p.Options = data

	return p
}

// 设置 Select 的模式为多选或标签，multiple | tags
func (p *Select) SetMode(mode string) *Select {
	p.Mode = mode

	return p
}

// 控件大小。注：标准表单内的输入框大小限制为 large。可选 large default small
func (p *Select) SetSize(size string) *Select {
	p.Size = size

	return p
}

// 可以点击清除图标删除内容
func (p *Select) SetAllowClear(allowClear bool) *Select {
	p.AllowClear = allowClear

	return p
}
