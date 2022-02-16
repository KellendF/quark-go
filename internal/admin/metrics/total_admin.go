package metrics

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/models"
	"github.com/quarkcms/quark-go/pkg/framework/db"
	"github.com/quarkcms/quark-go/pkg/ui/admin/metrics"
)

type TotalAdmin struct {
	metrics.Value
}

// 初始化
func (p *TotalAdmin) Init() {
	p.Title = "管理员数量"
	p.Col = 6
}

// 计算数值
func (p *TotalAdmin) Calculate(c *fiber.Ctx) interface{} {

	return p.Count(c, (&db.Model{}).Model(&models.Admin{})).SetValueStyle(map[string]string{"color": "#3f8600"})
}
