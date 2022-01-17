package page

type Component struct {
	Component string `json:"component"`
	Title     string `json:"title"`
	Body      string `json:"body"`
}

// 标题
func (p *Component) SetTitle(title string) *Component {
	p.Title = title
	return p
}

// 内容
func (p *Component) SetBody(body string) *Component {
	p.Body = body
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "page"

	return p
}
