package resources

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/admin/actions"
	"github.com/quarkcms/quark-go/internal/admin/searches"
	"github.com/quarkcms/quark-go/internal/models"
	"github.com/quarkcms/quark-go/pkg/framework/db"
	"github.com/quarkcms/quark-go/pkg/ui/admin"
)

type Admin struct {
	admin.Resource
}

// 初始化
func (p *Admin) Init() interface{} {

	// 标题
	p.Title = "管理员"

	// 模型
	p.Model = (&db.Model{}).Model(&models.Admin{})

	// 分页
	p.PerPage = 10

	return p
}

// 字段
func (p *Admin) Fields(c *fiber.Ctx) interface{} {
	field := &admin.Field{}

	return []interface{}{
		field.Text("username", "用户名"),
		field.Text("nickname", "昵称"),
	}
}

// 搜索
func (p *Admin) Searches(c *fiber.Ctx) interface{} {
	return []interface{}{
		(&searches.Input{}).Init("username", "用户名"),
	}
}

// 行为
func (p *Admin) Actions(c *fiber.Ctx) interface{} {
	return []interface{}{
		(&actions.CreateLink{}).Init(p.Title),
	}
}
