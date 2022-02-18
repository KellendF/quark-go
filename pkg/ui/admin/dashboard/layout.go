package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/models"
	"github.com/quarkcms/quark-go/pkg/ui/admin/utils"
	"github.com/quarkcms/quark-go/pkg/ui/component/layout"
	"github.com/quarkcms/quark-go/pkg/ui/component/page"
	"github.com/quarkcms/quark-go/pkg/ui/component/pagecontainer"
)

// 渲染页面组件
func PageComponentRender(c *fiber.Ctx, dashboard DashboardInterface, content interface{}) interface{} {
	layoutComponent := LayoutComponentRender(c, dashboard, content)

	return (&page.Component{}).
		SetBody(layoutComponent).
		JsonSerialize()
}

// 渲染页面布局组件
func LayoutComponentRender(c *fiber.Ctx, dashboard DashboardInterface, content interface{}) interface{} {

	// 获取登录管理员信息
	adminId := utils.Admin(c, "id")

	// 获取管理员菜单
	getMenus := (&models.Admin{}).GetMenus(adminId.(float64))

	return (&layout.Component{}).SetTitle("QuarkGo").
		SetBody(PageContainerComponentRender(content, dashboard)).
		SetContentWidth("Fluid").
		SetMenu(getMenus).
		SetFixSiderbar(true).
		SetFixedHeader(true).
		SetLayout("side").
		SetSiderWidth(208).
		SetLocale("zh-CN").
		SetNavTheme("dark").
		JsonSerialize()
}

// 渲染页面容器组件
func PageContainerComponentRender(content interface{}, dashboard DashboardInterface) interface{} {

	header := (&pagecontainer.PageHeader{}).Init().SetTitle(dashboard.GetTitle()).SetSubTitle(dashboard.GetSubTitle())

	return (&pagecontainer.Component{}).SetHeader(header).SetBody(content).JsonSerialize()
}

// 渲染列表页组件
func (p *Dashboard) DashboardComponentRender(c *fiber.Ctx, dashboard DashboardInterface) interface{} {

	return p.GetCards(c, dashboard)
}

// 渲染列表页组件
func (p *Dashboard) Render(c *fiber.Ctx, dashboard DashboardInterface, content interface{}) interface{} {

	// 初始化资源
	p.HandleInit(dashboard)

	return PageComponentRender(c, dashboard, content)
}
