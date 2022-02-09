package dashboards

import (
	"github.com/quarkcms/quark-go/pkg/ui/dashboard"
)

type Index struct {
	dashboard.Dashboard
}

// 初始化
func (dashboard *Index) Init() {

	dashboard.Title = "测试的标题"
}

// 内容
func (dashboard *Index) Cards() interface{} {

	return "测试内容"
}
