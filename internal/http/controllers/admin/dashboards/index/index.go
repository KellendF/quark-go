package index

import (
	"github.com/quarkcms/quark-go/pkg/dashboard"
)

type Component struct {
	dashboard.Dashboard
}

// 初始化
func (p *Component) Init() {

	p.Title = "测试的标题"
}
