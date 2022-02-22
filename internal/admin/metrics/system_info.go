package metrics

import (
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/admin/metrics"
	"github.com/quarkcms/quark-go/pkg/ui/component/descriptions"
)

type SystemInfo struct {
	metrics.Descriptions
}

// 初始化
func (p *SystemInfo) Init() *SystemInfo {
	p.Title = "系统信息"
	p.Col = 12

	return p
}

// 计算数值
func (p *SystemInfo) Calculate() *descriptions.Component {

	field := &descriptions.Field{}

	return p.Init().Result([]interface{}{
		field.Text("系统版本").SetValue("0.01"),
		field.Text("服务器操作系统").SetValue(runtime.GOOS),
		field.Text("Fiber版本").SetValue(fiber.Version),
	})
}
