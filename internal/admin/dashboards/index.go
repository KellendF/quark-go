package dashboards

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/admin/metrics"
	"github.com/quarkcms/quark-go/pkg/ui/admin"
)

type Index struct {
	admin.Dashboard
}

// 初始化
func (dashboard *Index) Init() {

	dashboard.Title = "仪表盘"
}

// 内容
func (dashboard *Index) Cards(c *fiber.Ctx) []any {
	return []any{
		&metrics.TotalAdmin{},
		&metrics.TotalLog{},
		&metrics.TotalPicture{},
		&metrics.TotalFile{},
		&metrics.SystemInfo{},
		&metrics.TeamInfo{},
	}
}
