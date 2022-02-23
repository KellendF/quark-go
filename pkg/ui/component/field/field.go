package field

import "github.com/quarkcms/quark-go/pkg/ui/component"

type Component struct {
	component.Element
	Title string `json:"title"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "field"

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// layout 的左上角 的 title
func (p *Component) SetTitle(title string) *Component {
	p.Title = title
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "field"

	return p
}
