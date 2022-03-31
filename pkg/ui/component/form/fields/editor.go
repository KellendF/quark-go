package fields

import "github.com/quarkcms/quark-go/pkg/ui/component"

type Editor struct {
	Item
}

// 初始化
func (p *Editor) Init() *Editor {
	p.Component = "editorField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.Style["height"] = 500
	p.Style["width"] = 800

	return p
}

// 高度
func (p *Editor) SetHeight(height int) *Editor {
	p.Style["height"] = height

	return p
}
