package admin

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/component/table"
	"gorm.io/gorm"
)

// 结构体
type Resource struct {
	Layout
	Title        string
	SubTitle     string
	PerPage      int
	IndexPolling int
	Model        *gorm.DB
}

// 获取模型
func (p *Resource) NewModel(resourceInstance interface{}) *gorm.DB {

	model := reflect.
		ValueOf(resourceInstance).
		Elem().
		FieldByName("Model").Interface()

	return model.(*gorm.DB)
}

// 列表标题
func (p *Resource) IndexTitle(c *fiber.Ctx, resourceInstance interface{}) string {
	return reflect.
		ValueOf(resourceInstance).
		Elem().
		FieldByName("Title").
		String() + "列表"
}

// 列表页表格主体
func (p *Resource) IndexExtraRender(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	return nil
}

// 列表页工具栏
func (p *Resource) IndexToolBar(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	return (&table.ToolBar{}).Init().SetTitle(p.IndexTitle(c, resourceInstance))
}

// 列表页组件渲染
func (p *Resource) IndexComponentRender(c *fiber.Ctx, resourceInstance interface{}, data interface{}) interface{} {

	// 列表标题
	title := p.IndexTitle(c, resourceInstance)

	// 反射获取参数
	value := reflect.ValueOf(resourceInstance).Elem()
	indexPolling := value.FieldByName("IndexPolling").Int()

	// 列表页表格主体
	indexExtraRender := p.IndexExtraRender(c, resourceInstance)

	// 列表页工具栏
	indexToolBar := p.IndexToolBar(c, resourceInstance)

	table := &table.Component{}
	component := table.
		Init().
		SetPolling(int(indexPolling)).
		SetTitle(title).
		SetTableExtraRender(indexExtraRender).
		SetToolBar(indexToolBar).
		SetColumns(p.IndexColumns(c, resourceInstance)).
		SetBatchActions("").
		SetSearches("").
		SetDatasource(data)

	return component
}
