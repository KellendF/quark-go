package metrics

import (
	"github.com/quarkcms/quark-go/pkg/ui/component/statistic"
)

type Descriptions struct {
	Metrics
	Precision int
}

// 包含组件的结果
func (p *Descriptions) Result(value int64) *statistic.Component {
	return (&statistic.Component{}).Init().SetTitle(p.Title).SetValue(value)
}
