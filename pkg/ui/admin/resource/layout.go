package resource

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/models"
	"github.com/quarkcms/quark-go/pkg/framework/config"
	"github.com/quarkcms/quark-go/pkg/ui/admin/utils"
	"github.com/quarkcms/quark-go/pkg/ui/component/footer"
	"github.com/quarkcms/quark-go/pkg/ui/component/layout"
	"github.com/quarkcms/quark-go/pkg/ui/component/page"
	"github.com/quarkcms/quark-go/pkg/ui/component/pagecontainer"
)

// 渲染页面组件
func PageComponentRender(c *fiber.Ctx, resource ResourceInterface, content interface{}) interface{} {
	layoutComponent := LayoutComponentRender(c, resource, content)

	return (&page.Component{}).
		SetBody(layoutComponent).
		JsonSerialize()
}

// 渲染页面布局组件
func LayoutComponentRender(c *fiber.Ctx, resource ResourceInterface, content interface{}) interface{} {

	// 获取登录管理员信息
	adminId := utils.Admin(c, "id")

	// 获取管理员菜单
	getMenus := (&models.Admin{}).GetMenus(adminId.(float64))

	// 页脚
	footer := (&footer.Component{}).
		Init().
		SetCopyright(config.Get("admin.copyright").(string)).
		SetLinks(config.Get("admin.links").([]map[string]interface{}))

	return (&layout.Component{}).
		Init().
		SetTitle(config.Get("admin.name").(string)).
		SetLogo(config.Get("admin.layout.logo")).
		SetHeaderActions(config.Get("admin.layout.header_actions").([]map[string]interface{})).
		SetLayout(config.Get("admin.layout.layout").(string)).
		SetSplitMenus(config.Get("admin.layout.split_menus").(bool)).
		SetHeaderTheme(config.Get("admin.layout.header_theme").(string)).
		SetContentWidth(config.Get("admin.layout.content_width").(string)).
		SetNavTheme(config.Get("admin.layout.nav_theme").(string)).
		SetPrimaryColor(config.Get("admin.layout.primary_color").(string)).
		SetFixSiderbar(config.Get("admin.layout.fix_siderbar").(bool)).
		SetFixedHeader(config.Get("admin.layout.fixed_header").(bool)).
		SetIconfontUrl(config.Get("admin.layout.iconfont_url").(string)).
		SetLocale(config.Get("admin.layout.locale").(string)).
		SetSiderWidth(config.Get("admin.layout.sider_width").(int)).
		SetMenu(getMenus).
		SetBody(PageContainerComponentRender(content, resource)).
		SetFooter(footer).
		JsonSerialize()
}

// 渲染页面容器组件
func PageContainerComponentRender(content interface{}, resource ResourceInterface) interface{} {
	component := &pagecontainer.Component{}
	return component.SetBody(content).JsonSerialize()
}

// 渲染列表页组件
func (p *Resource) IndexComponentRender() interface{} {
	return "xxxx"
}

// 渲染列表页组件
func (p *Resource) Render(c *fiber.Ctx, resource ResourceInterface, content interface{}) interface{} {

	// 初始化资源
	resource.HandleInit(resource)

	return PageComponentRender(c, resource, content)
}
