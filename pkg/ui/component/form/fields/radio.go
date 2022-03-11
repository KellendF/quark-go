package fields

import "github.com/quarkcms/quark-go/pkg/ui/component"

type Radio struct {
	Item
}

// 初始化
func (p *Radio) Init() *Radio {
	p.Component = "radioField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 设置单选属性
func (p *Radio) SetOptions(options map[string]interface{}) *Radio {
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

// 组件json序列化
func (p *Radio) JsonSerialize() *Radio {
	p.Component = "radioField"

	return p
}
