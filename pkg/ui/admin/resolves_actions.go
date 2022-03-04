package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/component/action"
)

// 列表行为
func (p *Resource) IndexActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	actions := resourceInstance.(interface{ Actions(*fiber.Ctx) interface{} }).Actions(c)

	var items []interface{}
	for _, v := range actions.([]interface{}) {
		shownOnIndex := v.(interface {
			ShownOnIndex() bool
		}).ShownOnIndex()

		if shownOnIndex {
			getAction := p.buildAction(v)
			items = append(items, getAction)
		}
	}

	return items
}

//表格行内行为
func (p *Resource) IndexTableRowActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	return ""
}

//表格多选弹出层行为
func (p *Resource) IndexTableAlertActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	return ""
}

//表单页行为
func (p *Resource) FormActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	return ""
}

//表单页右上角自定义区域行为
func (p *Resource) FormExtraActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	return ""
}

//详情页行为
func (p *Resource) DetailActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	return ""
}

//详情页右上角自定义区域行为
func (p *Resource) DetailExtraActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	return ""
}

//创建行为组件
func (p *Resource) buildAction(item interface{}) interface{} {
	name := item.(interface{ GetName() string }).GetName()
	withLoading := item.(interface{ GetWithLoading() bool }).GetWithLoading()

	action := (&action.Component{}).Init().SetLabel(name).SetWithLoading(withLoading)

	return action
}
