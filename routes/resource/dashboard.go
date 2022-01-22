package resource

import (
	"github.com/quarkcms/quark-go/internal/http/controllers/admin/dashboards/index"
	"github.com/quarkcms/quark-go/pkg/dashboard"
)

// 仪表盘资源
func Dashboard() map[string]dashboard.DashboardInterface {

	return map[string]dashboard.DashboardInterface{
		"index": &index.Component{},
	}
}
