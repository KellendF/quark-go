package descriptions

type Field struct {
	Component string `json:"component"`
}

// 初始化
func (p *Field) Init() *Field {
	p.Component = "descriptionField"

	return p
}

// 标题栏旁的头像
func (p *Field) text(label string) *Field {
	p.label = label
	return p
}

// 组件json序列化
func (p *Field) JsonSerialize() *Field {
	p.Component = "descriptionField"

	return p
}
