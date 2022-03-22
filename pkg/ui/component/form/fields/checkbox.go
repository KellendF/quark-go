package fields

import "github.com/quarkcms/quark-go/pkg/ui/component"

type Checkbox struct {
	Item
	Layout string
}

// 初始化
func (p *Checkbox) Init() *Checkbox {
	p.Component = "checkboxField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 设置单选属性
func (p *Checkbox) SetOptions(options map[interface{}]interface{}) *Checkbox {
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

// 配置 checkbox 的样子，支持垂直vertical 和 horizontal
func (p *Checkbox) SetLayout(layout string) *Checkbox {
	p.Layout = layout

	return p
}

// 组件json序列化
func (p *Checkbox) JsonSerialize() *Checkbox {
	p.Component = "checkboxField"

	return p
}
