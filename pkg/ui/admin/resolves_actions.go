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
	actions := resourceInstance.(interface{ Actions(*fiber.Ctx) interface{} }).Actions(c)

	var items []interface{}
	for _, v := range actions.([]interface{}) {
		shownOnIndexTableRow := v.(interface {
			ShownOnIndexTableRow() bool
		}).ShownOnIndexTableRow()

		if shownOnIndexTableRow {
			getAction := p.buildAction(v)
			items = append(items, getAction)
		}
	}

	return items
}

//表格多选弹出层行为
func (p *Resource) IndexTableAlertActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	actions := resourceInstance.(interface{ Actions(*fiber.Ctx) interface{} }).Actions(c)

	var items []interface{}
	for _, v := range actions.([]interface{}) {
		shownOnIndexTableAlert := v.(interface {
			ShownOnIndexTableAlert() bool
		}).ShownOnIndexTableAlert()

		if shownOnIndexTableAlert {
			getAction := p.buildAction(v)
			items = append(items, getAction)
		}
	}

	return items
}

//表单页行为
func (p *Resource) FormActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	actions := resourceInstance.(interface{ Actions(*fiber.Ctx) interface{} }).Actions(c)

	var items []interface{}
	for _, v := range actions.([]interface{}) {
		shownOnForm := v.(interface {
			ShownOnForm() bool
		}).ShownOnForm()

		if shownOnForm {
			getAction := p.buildAction(v)
			items = append(items, getAction)
		}
	}

	return items
}

//表单页右上角自定义区域行为
func (p *Resource) FormExtraActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	actions := resourceInstance.(interface{ Actions(*fiber.Ctx) interface{} }).Actions(c)

	var items []interface{}
	for _, v := range actions.([]interface{}) {
		shownOnFormExtra := v.(interface {
			ShownOnFormExtra() bool
		}).ShownOnFormExtra()

		if shownOnFormExtra {
			getAction := p.buildAction(v)
			items = append(items, getAction)
		}
	}

	return items
}

//详情页行为
func (p *Resource) DetailActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	actions := resourceInstance.(interface{ Actions(*fiber.Ctx) interface{} }).Actions(c)

	var items []interface{}
	for _, v := range actions.([]interface{}) {
		shownOnDetail := v.(interface {
			ShownOnDetail() bool
		}).ShownOnDetail()

		if shownOnDetail {
			getAction := p.buildAction(v)
			items = append(items, getAction)
		}
	}

	return items
}

//详情页右上角自定义区域行为
func (p *Resource) DetailExtraActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	actions := resourceInstance.(interface{ Actions(*fiber.Ctx) interface{} }).Actions(c)

	var items []interface{}
	for _, v := range actions.([]interface{}) {
		shownOnDetailExtra := v.(interface {
			ShownOnDetailExtra() bool
		}).ShownOnDetailExtra()

		if shownOnDetailExtra {
			getAction := p.buildAction(v)
			items = append(items, getAction)
		}
	}

	return items
}

//创建行为组件
func (p *Resource) buildAction(item interface{}) interface{} {
	name := item.(interface{ GetName() string }).GetName()
	withLoading := item.(interface{ GetWithLoading() bool }).GetWithLoading()
	reload := item.(interface{ GetReload() string }).GetReload()
	api := item.(interface{ GetApi() string }).GetApi()
	actionType := item.(interface{ GetActionType() string }).GetActionType()
	buttonType := item.(interface{ GetType() string }).GetType()
	size := item.(interface{ GetSize() string }).GetSize()
	icon := item.(interface{ GetIcon() string }).GetIcon()
	confirmTitle := item.(interface{ GetConfirmTitle() string }).GetConfirmTitle()
	confirmText := item.(interface{ GetConfirmText() string }).GetConfirmText()
	confirmType := item.(interface{ GetConfirmType() string }).GetConfirmType()

	action := (&action.Component{}).
		Init().
		SetLabel(name).
		SetWithLoading(withLoading).
		SetReload(reload).SetApi(api).
		SetActionType(actionType).
		SetType(buttonType, false).
		SetSize(size)

	if icon != "" {
		action = action.
			SetIcon(icon)
	}

	switch actionType {
	case "link":
		href := item.(interface{ GetHref() string }).GetHref()
		target := item.(interface{ GetTarget() string }).GetTarget()

		action = action.
			SetLink(href, target)
	case "modal":
		// todo
	case "drawer":
		// todo
	}

	if confirmTitle != "" {
		action = action.
			SetWithConfirm(confirmTitle, confirmText, confirmType)
	}

	return action
}
