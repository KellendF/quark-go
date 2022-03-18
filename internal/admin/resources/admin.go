package resources

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/admin/actions"
	"github.com/quarkcms/quark-go/internal/admin/searches"
	"github.com/quarkcms/quark-go/internal/models"
	"github.com/quarkcms/quark-go/pkg/framework/db"
	"github.com/quarkcms/quark-go/pkg/ui/admin"
	"github.com/quarkcms/quark-go/pkg/ui/component/table"
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
		field.Text("id", "ID"),
		field.Text("username", "用户名", func() interface{} {
			id := strconv.Itoa(p.Field["id"].(int))
			username := p.Field["username"].(string)

			return "<a href='#/index?api=admin/admin/edit&id=" + id + "'>" + username + "</a>"
		}),
		field.Text("nickname", "昵称").SetEditable(true),
		field.Text("email", "邮箱"),
		field.Text("phone", "手机号"),
		field.Radio("sex", "性别").
			SetOptions(map[string]interface{}{
				"1": "男",
				"2": "女",
			}).SetDefault("1").
			SetColumn(func(column *table.Column) *table.Column {
				// return column.SetSorter(true)
				return column.SetFilters(true)
			}),
		field.Datetime("last_login_time", "最后登录时间"),
		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true).
			SetDefault(true),
	}
}

// 搜索
func (p *Admin) Searches(c *fiber.Ctx) []interface{} {
	return []interface{}{
		(&searches.Input{}).Init("username", "用户名"),
		(&searches.Input{}).Init("nickname", "昵称"),
		(&searches.Status{}).Init(),
		(&searches.DateTimeRange{}).Init("last_login_time", "登录时间"),
	}
}

// 行为
func (p *Admin) Actions(c *fiber.Ctx) interface{} {
	return []interface{}{
		(&actions.CreateLink{}).Init(p.Title),
		(&actions.Delete{}).Init("批量删除").SetOnlyOnIndexTableAlert(true),
		(&actions.Disable{}).Init("批量禁用"),
		(&actions.Enable{}).Init("批量启用"),
		(&actions.ChangeStatus{}).Init(),
		(&actions.EditLink{}).Init("编辑"),
		(&actions.Delete{}).Init("删除").SetOnlyOnIndexTableRow(true),
		(&actions.FormSubmit{}).Init(),
		(&actions.FormReset{}).Init(),
		(&actions.FormBack{}).Init(),
		(&actions.FormExtraBack{}).Init(),
	}
}
