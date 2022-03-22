package resources

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/admin/actions"
	"github.com/quarkcms/quark-go/internal/admin/searches"
	"github.com/quarkcms/quark-go/internal/models"
	"github.com/quarkcms/quark-go/pkg/framework/db"
	"github.com/quarkcms/quark-go/pkg/ui/admin"
)

type Article struct {
	admin.Resource
}

// 初始化
func (p *Article) Init() interface{} {

	// 标题
	p.Title = "文章"

	// 模型
	p.Model = (&db.Model{}).Model(&models.Post{})

	// 分页
	p.PerPage = 10

	return p
}

// 字段
func (p *Article) Fields(c *fiber.Ctx) interface{} {
	field := &admin.Field{}

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Text("title", "标题"),
		field.Datetime("created_at", "创建时间").OnlyOnIndex(),
		field.Datetime("updated_at", "更新时间").OnlyOnIndex(),
	}
}

// 搜索
func (p *Article) Searches(c *fiber.Ctx) []interface{} {
	return []interface{}{
		(&searches.Input{}).Init("name", "名称"),
	}
}

// 行为
func (p *Article) Actions(c *fiber.Ctx) interface{} {
	return []interface{}{
		(&actions.CreateLink{}).Init(p.Title),
		(&actions.Delete{}).Init("批量删除"),
		(&actions.EditLink{}).Init("编辑"),
		(&actions.Delete{}).Init("删除"),
		(&actions.FormSubmit{}).Init(),
		(&actions.FormReset{}).Init(),
		(&actions.FormBack{}).Init(),
		(&actions.FormExtraBack{}).Init(),
	}
}
