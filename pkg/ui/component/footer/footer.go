package footer

type Component struct {
	Component string                   `json:"component"`
	Copyright string                   `json:"copyright"`
	Links     []map[string]interface{} `json:"links"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "footer"

	return p
}

// 版权信息
func (p *Component) SetCopyright(copyright string) *Component {
	p.Copyright = copyright
	return p
}

// 版权信息
func (p *Component) SetLinks(links []map[string]interface{}) *Component {
	p.Links = links
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "footer"

	return p
}
