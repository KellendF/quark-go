package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/admin/actions"
	"github.com/quarkcms/quark-go/pkg/ui/admin/component/action"
	"github.com/quarkcms/quark-go/pkg/ui/admin/component/form"
)

type EditModal struct {
	actions.Modal
}

// 初始化
func (p *EditModal) Init(name string) *EditModal {
	// 初始化父结构
	p.ParentInit()

	// 类型
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 文字
	p.Name = name

	// 关闭时销毁 Modal 里的子元素
	p.DestroyOnClose = true

	// 执行成功后刷新的组件
	p.Reload = "table"

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	return p
}

// 内容
func (p *EditModal) GetBody(c *fiber.Ctx, resourceInstance interface{}) interface{} {

	api := resourceInstance.(interface {
		UpdateApi(*fiber.Ctx, interface{}) string
	}).UpdateApi(c, resourceInstance)

	initApi := resourceInstance.(interface {
		EditValueApi(*fiber.Ctx) string
	}).EditValueApi(c)

	fields := resourceInstance.(interface {
		UpdateFieldsWithinComponents(*fiber.Ctx, interface{}) interface{}
	}).UpdateFieldsWithinComponents(c, resourceInstance)

	return (&form.Component{}).
		Init().
		SetKey("editModalForm", false).
		SetApi(api).
		SetInitApi(initApi).
		SetBody(fields).
		SetLabelCol(map[string]interface{}{
			"span": 6,
		}).
		SetWrapperCol(map[string]interface{}{
			"span": 18,
		})
}

// 弹窗行为
func (p *EditModal) GetActions(c *fiber.Ctx, resourceInstance interface{}) []interface{} {

	return []interface{}{
		(&action.Component{}).
			Init().
			SetLabel("取消").
			SetActionType("cancel"),

		(&action.Component{}).
			Init().
			SetLabel("提交").
			SetWithLoading(true).
			SetReload("table").
			SetActionType("submit").
			SetType("primary", false).
			SetSubmitForm("editModalForm"),
	}
}
