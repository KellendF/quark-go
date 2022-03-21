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
		field.Hidden("id", "ID"),

		field.Text("username", "用户名", func() interface{} {

			return "<a href='#/index?api=admin/admin/edit&id=" + strconv.Itoa(p.Field["id"].(int)) + "'>" + p.Field["username"].(string) + "</a>"
		}).
			SetEditable(true).
			SetRules(
				[]string{
					"required",
					"min:6",
					"max:20",
				},
				map[string]string{
					"required": "用户名必须填写",
					"min":      "用户名不能少于6个字符",
					"max":      "用户名不能超过20个字符",
				},
			).
			SetCreationRules(
				[]string{
					"unique:admins,username",
				},
				map[string]string{
					"unique": "用户名已存在",
				},
			).
			SetUpdateRules(
				[]string{
					"unique:admins,username,{id}",
				},
				map[string]string{
					"unique": "用户名已存在",
				},
			),

		field.Text("nickname", "昵称").
			SetEditable(true).
			SetRules(
				[]string{
					"required",
				},
				map[string]string{
					"required": "昵称必须填写",
				},
			),

		field.Text("email", "邮箱").
			SetRules(
				[]string{
					"required",
				},
				map[string]string{
					"required": "邮箱必须填写",
				},
			).
			SetCreationRules(
				[]string{
					"unique:admins,email",
				},
				map[string]string{
					"unique": "邮箱已存在",
				},
			).
			SetUpdateRules(
				[]string{
					"unique:admins,email,{id}",
				},
				map[string]string{
					"unique": "邮箱已存在",
				},
			),

		field.Text("phone", "手机号").
			SetRules(
				[]string{
					"required",
				},
				map[string]string{
					"required": "手机号必须填写",
				},
			).
			SetCreationRules(
				[]string{
					"unique:admins,phone",
				},
				map[string]string{
					"unique": "手机号已存在",
				},
			).
			SetUpdateRules(
				[]string{
					"unique:admins,phone,{id}",
				},
				map[string]string{
					"unique": "手机号已存在",
				},
			),

		field.Radio("sex", "性别").
			SetOptions(map[string]interface{}{
				"1": "男",
				"2": "女",
			}).SetDefault("1").
			SetColumn(func(column *table.Column) *table.Column {
				// return column.SetSorter(true)
				return column.SetFilters(true)
			}),

		field.Text("password", "密码").
			SetCreationRules(
				[]string{
					"required",
				},
				map[string]string{
					"required": "密码必须填写",
				},
			).OnlyOnForms(),

		field.Datetime("last_login_time", "最后登录时间").OnlyOnIndex(),

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
		(&actions.Delete{}).Init("批量删除"),
		(&actions.Disable{}).Init("批量禁用"),
		(&actions.Enable{}).Init("批量启用"),
		(&actions.ChangeStatus{}).Init(),
		(&actions.EditLink{}).Init("编辑"),
		(&actions.Delete{}).Init("删除"),
		(&actions.FormSubmit{}).Init(),
		(&actions.FormReset{}).Init(),
		(&actions.FormBack{}).Init(),
		(&actions.FormExtraBack{}).Init(),
	}
}
