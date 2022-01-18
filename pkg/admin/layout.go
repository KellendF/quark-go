package admin

import (
	"github.com/quarkcms/quark-go/pkg/component/layout"
	"github.com/quarkcms/quark-go/pkg/component/page"
	"github.com/quarkcms/quark-go/pkg/component/pagecontainer"
)

// 渲染页面组件
func PageComponentRender(content interface{}) interface{} {
	component := &page.Component{}

	layoutComponent := LayoutComponentRender(content)

	p := &Resource{}

	return component.
		SetTitle(p.Title).
		SetBody(layoutComponent).
		JsonSerialize()
}

// 渲染页面布局组件
func LayoutComponentRender(content interface{}) interface{} {

	p := &Resource{}

	component := &layout.Component{}
	return component.SetTitle(p.Title).
		SetBody(PageContainerComponentRender(content)).
		JsonSerialize()
}

// 渲染页面容器组件
func PageContainerComponentRender(content interface{}) interface{} {
	component := &pagecontainer.Component{}
	return component.SetBody(content).JsonSerialize()
}

// 渲染列表页组件
func (p *Resource) IndexComponentRender() interface{} {
	return "xxx"
}

// 渲染列表页组件
func (p *Resource) Render(content interface{}) interface{} {
	return PageComponentRender(content)
}
