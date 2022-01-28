package action

type Component struct {
	Component string `json:"component"`
	Title     string `json:"title"`
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
