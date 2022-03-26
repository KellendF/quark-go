package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/admin/actions"
)

type CreateDrawer struct {
	actions.Drawer
}

// 初始化
func (p *CreateDrawer) Init(name string) *CreateDrawer {
	// 初始化父结构
	p.ParentInit()

	// 类型
	p.Type = "primary"

	// 图标
	p.Icon = "plus-circle"

	// 文字
	p.Name = "创建" + name

	// 执行成功后刷新的组件
	p.Reload = "table"

	// 设置展示位置
	p.SetOnlyOnIndex(true)

	return p
}

// 内容
func (p *CreateDrawer) GetBody(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	//  $request = new ResourceCreateRequest;

	//  // 表单
	//  return Form::key('createDrawerForm')
	//  ->api($request->newResource()->creationApi($request))
	//  ->items($request->newResource()->creationFieldsWithinComponents($request))
	//  ->initialValues($request->newResource()->beforeCreating($request))
	//  ->labelCol([
	// 	 'span' => 6
	//  ])
	//  ->wrapperCol([
	// 	 'span' => 18
	//  ]);

	return nil
}

// 弹窗行为
func (p *CreateDrawer) GetActions(c *fiber.Ctx, resourceInstance interface{}) []interface{} {
	//  return [
	// 	 Action::make('取消')->actionType('cancel'),

	// 	 Action::make("提交")
	// 	 ->reload('table')
	// 	 ->type('primary')
	// 	 ->actionType('submit')
	// 	 ->submitForm('createDrawerForm')
	//  ];

	return []interface{}{}
}
