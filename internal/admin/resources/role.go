package resources

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/admin/actions"
	"github.com/quarkcms/quark-go/internal/admin/searches"
	"github.com/quarkcms/quark-go/internal/models"
	"github.com/quarkcms/quark-go/pkg/framework/db"
	"github.com/quarkcms/quark-go/pkg/ui/admin"
)

type Role struct {
	admin.Resource
}

// 初始化
func (p *Role) Init() interface{} {

	// 标题
	p.Title = "角色"

	// 模型
	p.Model = (&db.Model{}).Model(&models.Role{})

	// 分页
	p.PerPage = 10

	return p
}

// 字段
func (p *Role) Fields(c *fiber.Ctx) interface{} {
	field := &admin.Field{}
	treeData := (&models.Menu{}).Tree()

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Text("name", "名称").
			SetRules(
				[]string{
					"required",
				},
				map[string]string{
					"required": "名称必须填写",
				},
			),

		field.Text("guard_name", "GuardName").SetDefault("admin"),
		field.Tree("menu_ids", "权限").SetData(treeData).OnlyOnForms(),
		field.Datetime("created_at", "创建时间").OnlyOnIndex(),
		field.Datetime("updated_at", "更新时间").OnlyOnIndex(),
	}
}

// 搜索
func (p *Role) Searches(c *fiber.Ctx) []interface{} {
	return []interface{}{
		(&searches.Input{}).Init("name", "名称"),
	}
}

// 行为
func (p *Role) Actions(c *fiber.Ctx) interface{} {
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

// 编辑页面显示前回调
func (p *Role) BeforeEditing(c *fiber.Ctx, data map[string]interface{}) map[string]interface{} {
	return data
}

// 保存数据前回调
func (p *Role) BeforeSaving(c *fiber.Ctx, submitData map[string]interface{}) interface{} {
	return submitData
}
