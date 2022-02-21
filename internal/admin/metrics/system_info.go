package metrics

import (
	"github.com/quarkcms/quark-go/internal/models"
	"github.com/quarkcms/quark-go/pkg/framework/db"
	"github.com/quarkcms/quark-go/pkg/ui/admin/metrics"
	"github.com/quarkcms/quark-go/pkg/ui/component/statistic"
)

type SystemInfo struct {
	metrics.Value
}

// 初始化
func (p *SystemInfo) Init() *SystemInfo {
	p.Title = "团队信息"
	p.Col = 12

	return p
}

// 计算数值
func (p *SystemInfo) Calculate() *statistic.Component {

	return p.
		Init().
		Count((&db.Model{}).Model(&models.Admin{})).
		SetValueStyle(map[string]string{"color": "#3f8600"})
}
