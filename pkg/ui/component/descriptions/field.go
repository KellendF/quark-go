package descriptions

import "github.com/quarkcms/quark-go/pkg/ui/component/descriptions/fields"

type Field struct {
	Component string `json:"component"`
}

// 初始化
func (p *Field) Init() *Field {
	p.Component = "descriptionField"

	return p
}

// text组件
func (p *Field) Text(params ...string) *fields.Text {

	fields := &fields.Text{}

	if len(params) == 1 {
		fields = fields.Init().SetDataIndex(params[0]).SetLabel(params[0])
	} else {
		fields = fields.Init().SetDataIndex(params[0]).SetLabel(params[1])
	}

	return fields
}
