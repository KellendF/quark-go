package index

import (
	"github.com/quarkcms/quark-go/pkg/ui/dashboard"
)

type Component struct {
	dashboard.Dashboard
}

// 初始化
func (p *Component) Init() {

	p.Title = "测试的标题"
}

// 内容
func (p *Component) Cards() interface{} {

	return "测试内容"
}
