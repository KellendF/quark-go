package resource

import (
	"github.com/quarkcms/quark-go/internal/http/controllers/admin/dashboards"
	"github.com/quarkcms/quark-go/pkg/ui/dashboard"
)

// 仪表盘资源
func Dashboard() map[string]dashboard.DashboardInterface {

	return map[string]dashboard.DashboardInterface{
		"index": &dashboards.Index{},
	}
}
