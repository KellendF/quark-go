package admin

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/component/action"
	"github.com/quarkcms/quark-go/pkg/ui/component/space"
)

// 列表行为
func (p *Resource) IndexActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	actions := resourceInstance.(interface {
		Actions(*fiber.Ctx) []interface{}
	}).Actions(c)

	var items []interface{}
	for _, v := range actions {
		shownOnIndex := v.(interface {
			ShownOnIndex() bool
		}).ShownOnIndex()

		if shownOnIndex {
			getAction := p.buildAction(c, v, resourceInstance)
			items = append(items, getAction)
		}
	}

	return (&space.Component{}).Init().SetBody(items)
}

//表格行内行为
func (p *Resource) IndexTableRowActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	actions := resourceInstance.(interface {
		Actions(*fiber.Ctx) []interface{}
	}).Actions(c)

	var items []interface{}
	for _, v := range actions {
		shownOnIndexTableRow := v.(interface {
			ShownOnIndexTableRow() bool
		}).ShownOnIndexTableRow()

		if shownOnIndexTableRow {
			getAction := p.buildAction(c, v, resourceInstance)
			items = append(items, getAction)
		}
	}

	return items
}

//表格多选弹出层行为
func (p *Resource) IndexTableAlertActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	actions := resourceInstance.(interface {
		Actions(*fiber.Ctx) []interface{}
	}).Actions(c)

	var items []interface{}
	for _, v := range actions {
		shownOnIndexTableAlert := v.(interface {
			ShownOnIndexTableAlert() bool
		}).ShownOnIndexTableAlert()

		if shownOnIndexTableAlert {
			getAction := p.buildAction(c, v, resourceInstance)
			items = append(items, getAction)
		}
	}

	return items
}

//表单页行为
func (p *Resource) FormActions(c *fiber.Ctx, resourceInstance interface{}) []interface{} {
	actions := resourceInstance.(interface {
		Actions(*fiber.Ctx) []interface{}
	}).Actions(c)

	var items []interface{}
	for _, v := range actions {
		shownOnForm := v.(interface {
			ShownOnForm() bool
		}).ShownOnForm()

		if shownOnForm {
			getAction := p.buildAction(c, v, resourceInstance)
			items = append(items, getAction)
		}
	}

	return items
}

//表单页右上角自定义区域行为
func (p *Resource) FormExtraActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	actions := resourceInstance.(interface {
		Actions(*fiber.Ctx) []interface{}
	}).Actions(c)

	var items []interface{}
	for _, v := range actions {
		shownOnFormExtra := v.(interface {
			ShownOnFormExtra() bool
		}).ShownOnFormExtra()

		if shownOnFormExtra {
			getAction := p.buildAction(c, v, resourceInstance)
			items = append(items, getAction)
		}
	}

	return items
}

//详情页行为
func (p *Resource) DetailActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	actions := resourceInstance.(interface {
		Actions(*fiber.Ctx) []interface{}
	}).Actions(c)

	var items []interface{}
	for _, v := range actions {
		shownOnDetail := v.(interface {
			ShownOnDetail() bool
		}).ShownOnDetail()

		if shownOnDetail {
			getAction := p.buildAction(c, v, resourceInstance)
			items = append(items, getAction)
		}
	}

	return items
}

//详情页右上角自定义区域行为
func (p *Resource) DetailExtraActions(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	actions := resourceInstance.(interface {
		Actions(*fiber.Ctx) []interface{}
	}).Actions(c)

	var items []interface{}
	for _, v := range actions {
		shownOnDetailExtra := v.(interface {
			ShownOnDetailExtra() bool
		}).ShownOnDetailExtra()

		if shownOnDetailExtra {
			getAction := p.buildAction(c, v, resourceInstance)
			items = append(items, getAction)
		}
	}

	return items
}

//创建行为组件
func (p *Resource) buildAction(c *fiber.Ctx, item interface{}, resourceInstance interface{}) interface{} {
	name := item.(interface{ GetName() string }).GetName()
	withLoading := item.(interface{ GetWithLoading() bool }).GetWithLoading()
	reload := item.(interface{ GetReload() string }).GetReload()

	// uri唯一标识
	uriKey := item.(interface {
		GetUriKey(interface{}) string
	}).GetUriKey(item)

	// 获取api
	api := item.(interface {
		GetApi(*fiber.Ctx) string
	}).GetApi(c)

	// 获取api替换参数
	params := item.(interface {
		GetApiParams() []string
	}).GetApiParams()

	if api == "" {
		api = p.buildActionApi(c, params, uriKey)
	}

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
		SetReload(reload).
		SetApi(api).
		SetActionType(actionType).
		SetType(buttonType, false).
		SetSize(size)

	if icon != "" {
		action = action.
			SetIcon(icon)
	}

	switch actionType {
	case "link":
		href := item.(interface{ GetHref(c *fiber.Ctx) string }).GetHref(c)
		target := item.(interface{ GetTarget(c *fiber.Ctx) string }).GetTarget(c)

		action = action.
			SetLink(href, target)
	case "modal":
		// todo
	case "drawer":
		// formBody := item.(interface{ GetBody(c *fiber.Ctx,resourceInstance interface{}) interface{} }).GetBody(c,resourceInstance)
		// formActions := item.(interface{ GetActions(c *fiber.Ctx,resourceInstance interface{}) []interface{} }).GetActions(c,resourceInstance)

		// action = action.SetDrawer()
	}

	if confirmTitle != "" {
		action = action.
			SetWithConfirm(confirmTitle, confirmText, confirmType)
	}

	return action
}

//创建行为接口
func (p *Resource) buildActionApi(c *fiber.Ctx, params []string, uriKey string) string {
	paramsUri := ""

	for _, v := range params {
		paramsUri = paramsUri + v + "=${" + v + "}&"
	}

	api := strings.Replace(strings.Replace(c.Path(), "/api/", "", -1), "/index", "/action/"+uriKey, -1)
	if paramsUri != "" {
		api = api + "?" + paramsUri
	}

	return api
}
