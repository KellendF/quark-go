package dropdown

import "github.com/quarkcms/quark-go/pkg/ui/component"

type Component struct {
	component.Element
	Label    string `json:"label"`
	Block    bool   `json:"block"`
	Danger   bool   `json:"danger"`
	Disabled bool   `json:"disabled"`
	Ghost    bool   `json:"ghost"`
	Icon     string `json:"icon"`
	Shape    string `json:"shape"`
	Size     string `json:"size"`
	Type     string `json:"type"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "dropdown"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 设置按钮文字
func (p *Component) SetLabel(label string) *Component {
	p.Label = label

	return p
}

// 将按钮宽度调整为其父宽度的选项
func (p *Component) SetBlock(block bool) *Component {
	p.Block = block

	return p
}

// 设置危险按钮
func (p *Component) SetDanger(danger bool) *Component {
	p.Danger = danger

	return p
}

// 按钮失效状态
func (p *Component) SetDisabled(disabled bool) *Component {
	p.Disabled = disabled

	return p
}

// 幽灵属性，使按钮背景透明
func (p *Component) SetGhost(ghost bool) *Component {
	p.Ghost = ghost

	return p
}

// 设置按钮图标
func (p *Component) SetIcon(icon string) *Component {
	p.Icon = "icon-" + icon

	return p
}

// 设置按钮形状，可选值为 circle、 round 或者不设
func (p *Component) SetShape(shape string) *Component {
	p.Shape = shape

	return p
}

// 设置按钮类型，primary | ghost | dashed | link | text | default
func (p *Component) SetType(buttonType string, danger bool) *Component {
	p.Type = buttonType
	p.Danger = danger

	return p
}

// 设置按钮大小，large | middle | small | default
func (p *Component) SetSize(size string) *Component {
	p.Size = size

	return p
}
