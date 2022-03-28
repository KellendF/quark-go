package resources

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/admin/actions"
	"github.com/quarkcms/quark-go/internal/admin/searches"
	"github.com/quarkcms/quark-go/internal/models"
	"github.com/quarkcms/quark-go/pkg/ui/admin"
)

type Picture struct {
	admin.Resource
}

// 初始化
func (p *Picture) Init() interface{} {

	// 标题
	p.Title = "图片"

	// 模型
	p.Model = &models.Picture{}

	// 分页
	p.PerPage = 10

	return p
}

// 字段
func (p *Picture) Fields(c *fiber.Ctx) []interface{} {
	field := &admin.Field{}

	return []interface{}{
		field.ID("id", "ID"),
		field.Text("name", "名称"),
		field.Text("size", "大小"),
		field.Text("width", "宽度"),
		field.Text("height", "高度"),
		field.Text("ext", "扩展名"),
		field.Datetime("created_at", "上传时间"),
	}
}

// 搜索
func (p *Picture) Searches(c *fiber.Ctx) []interface{} {
	return []interface{}{
		(&searches.Input{}).Init("name", "名称"),
		(&searches.DateTimeRange{}).Init("created_at", "上传时间"),
	}
}

// 行为
func (p *Picture) Actions(c *fiber.Ctx) []interface{} {
	return []interface{}{
		(&actions.Delete{}).Init("批量删除"),
		(&actions.Delete{}).Init("删除"),
	}
}
