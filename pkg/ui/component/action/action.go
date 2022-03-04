package action

import "github.com/quarkcms/quark-go/pkg/ui/component"

type Component struct {
	component.Element
	Label        string      `json:"label"`
	Block        bool        `json:"block"`
	Danger       bool        `json:"danger"`
	Disabled     bool        `json:"disabled"`
	Ghost        bool        `json:"ghost"`
	Icon         string      `json:"icon"`
	Shape        string      `json:"shape"`
	Size         string      `json:"size"`
	Type         string      `json:"type"`
	ActionType   string      `json:"actionType"`
	SubmitForm   string      `json:"submitForm"`
	Href         string      `json:"href"`
	Target       string      `json:"target"`
	Modal        interface{} `json:"modal"`
	Drawer       interface{} `json:"drawer"`
	ConfirmTitle string      `json:"confirmTitle"`
	ConfirmText  string      `json:"confirmText"`
	ConfirmType  string      `json:"confirmType"`
	Api          string      `json:"api"`
	Reload       string      `json:"reload"`
	WithLoading  bool        `json:"withLoading"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "action"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

/**
 * 将按钮宽度调整为其父宽度的选项
 *
 * @param  bool  block
 * @return p
 */
func (p *Component) SetBlock(block bool) *Component {
	p.Block = block

	return p
}

/**
 * 设置危险按钮
 *
 * @param  bool  danger
 * @return p
 */
func (p *Component) SetDanger(danger bool) *Component {
	p.Danger = danger

	return p
}

/**
 * 按钮失效状态
 *
 * @param  bool  disabled
 * @return p
 */
func (p *Component) SetDisabled(disabled bool) *Component {
	p.Disabled = disabled

	return p
}

/**
 * 幽灵属性，使按钮背景透明
 *
 * @param  bool  ghost
 * @return p
 */
func (p *Component) SetGhost(ghost bool) *Component {
	p.Ghost = ghost

	return p
}

/**
 * 设置按钮图标
 *
 * @param  string  icon
 * @return p
 */
func (p *Component) SetIcon(icon string) *Component {
	p.Icon = "icon-" + icon

	return p
}

/**
 * 设置按钮形状，可选值为 circle、 round 或者不设
 *
 * @param  string  shape
 * @return p
 */
func (p *Component) SetShape(shape string) *Component {
	p.Shape = shape

	return p
}

/**
 * 设置按钮类型，primary | ghost | dashed | link | text | default
 *
 * @param  string  type
 * @param  bool  danger
 * @return p
 */
func (p *Component) SetType(buttonType string, danger bool) *Component {
	p.Type = buttonType
	p.Danger = danger

	return p
}

/**
 * 设置按钮大小，large | middle | small | default
 *
 * @param  string  type
 * @param  bool  danger
 * @return p
 */
func (p *Component) SetSize(size string) *Component {
	p.Size = size

	return p
}

/**
 * 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
 *
 * @param  string  actionType
 * @return p
 */
func (p *Component) SetActionType(actionType string) *Component {
	p.ActionType = actionType

	return p
}

/**
 * 当action 的作用类型为submit的时候，可以指定提交哪个表格，submitForm为提交表单的key值，为空时提交当前表单
 *
 * @param  string  formKey
 * @return p
 */
func (p *Component) SetSubmitForm(formKey string) *Component {
	p.SubmitForm = formKey

	return p
}

/**
 * 点击跳转的地址，指定此属性 button 的行为和 a 链接一致
 *
 * @param  string  href
 * @return p
 */
func (p *Component) SetHref(href string) *Component {
	p.Href = href

	return p
}

/**
 * 相当于 a 链接的 target 属性，href 存在时生效
 *
 * @param  string  target
 * @return p
 */
func (p *Component) SetTarget(target string) *Component {
	p.Target = target

	return p
}

/**
 * 设置跳转链接
 *
 * @param  string  href
 * @param  string  target _blank,_self,_parent
 * @return p
 */
func (p *Component) SetLink(href string, target string) *Component {
	p.SetHref(href)
	p.SetTarget(target)
	p.ActionType = "link"

	return p
}

/**
 * 弹窗
 *
 * @param  Closure  modal
 * @return p
 */
func (p *Component) SetModal(callback interface{}) *Component {
	// todo
	return p
}

/**
 * 抽屉
 *
 * @param  Closure  drawer
 * @return p
 */
func (p *Component) SetDrawer(callback interface{}) *Component {
	// todo
	return p
}

/**
 * 设置行为前的确认操作
 *
 * @param  string  title
 * @param  string  text
 * @param  string  type
 * @return p
 */
func (p *Component) SetWithConfirm(title string, text string, confirmType string) *Component {
	p.ConfirmTitle = title
	p.ConfirmText = text
	p.ConfirmType = confirmType

	return p
}

/**
 *  执行行为的接口链接
 *
 * @param  string  api
 * @return p
 */
func (p *Component) SetApi(api string) *Component {
	p.Api = api
	p.ActionType = "ajax"

	return p
}

/**
 *  执行成功后刷新的组件
 *
 * @param  string  reload
 * @return p
 */
func (p *Component) SetReload(reload string) *Component {
	p.Reload = reload

	return p
}

/**
 *  是否具有loading
 *
 * @param  bool  loading
 * @return p
 */
func (p *Component) SetWithLoading(loading bool) *Component {
	p.WithLoading = loading

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "action"

	return p
}
