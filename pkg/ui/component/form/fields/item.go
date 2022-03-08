package fields

import (
	"github.com/quarkcms/quark-go/pkg/ui/component"
)

type Item struct {
	component.Element
	Tooltip              string      `json:"tooltip"`
	Width                int         `json:"width"`
	Colon                bool        `json:"colon"`
	Value                interface{} `json:"value"`
	DefaultValue         interface{} `json:"defaultValue"`
	Extra                string      `json:"extra"`
	HasFeedback          bool        `json:"hasFeedback"`
	Help                 string      `json:"help"`
	NoStyle              bool        `json:"noStyle"`
	Label                string      `json:"label"`
	LabelAlign           string      `json:"labelAlign"`
	LabelCol             interface{} `json:"labelCol"`
	Name                 string      `json:"name"`
	Required             bool        `json:"required"`
	Disabled             bool        `json:"disabled"`
	Ignore               bool        `json:"ignore"`
	Rules                interface{} `json:"rules"`
	RuleMessages         interface{} `json:"ruleMessages"`
	CreationRules        interface{} `json:"creationRules"`
	CreationRuleMessages interface{} `json:"creationRuleMessages"`
	UpdateRules          interface{} `json:"updateRules"`
	UpdateRuleMessages   interface{} `json:"updateRuleMessages"`
	FrontendRules        interface{} `json:"frontendRules"`
	ValuePropName        string      `json:"valuePropName"`
	WrapperCol           interface{} `json:"wrapperCol"`
	When                 interface{} `json:"when"`
	ShowOnIndex          bool        `json:"showOnIndex"`
	ShowOnDetail         bool        `json:"showOnDetail"`
	ShowOnCreation       bool        `json:"showOnCreation"`
	ShowOnUpdate         bool        `json:"showOnUpdate"`
	ShowOnExport         bool        `json:"showOnExport"`
	ShowOnImport         bool        `json:"showOnImport"`
	Editable             bool        `json:"editable"`
	Column               interface{} `json:"column"`
	Options              interface{} `json:"options"`
}

// 初始化
func (p *Item) InitItem() *Item {
	p.Colon = true
	p.LabelAlign = "right"
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = true
	p.ShowOnImport = true

	return p
}

// Set style.
func (p *Item) SetStyle(style map[string]interface{}) *Item {
	p.Style = style

	return p
}

//会在 label 旁增加一个 icon，悬浮后展示配置的信息
func (p *Item) SetTooltip(tooltip string) *Item {
	p.Tooltip = tooltip

	return p
}

/**
 * Field 的长度，我们归纳了常用的 Field 长度以及适合的场景，支持了一些枚举 "xs" , "s" , "m" , "l" , "x"
 *
 * @param  string|number width
 * @return p
 */
func (p *Item) SetWidth(width int) *Item {
	// if(!in_array(width,["xs","s","m","l","x"])) {
	//     throw new Exception("argument must be "xs","s","m","l","x"!")
	// }

	p.Style["width"] = width
	return p
}

/**
 * 配合 label 属性使用，表示是否显示 label 后面的冒号
 *
 * @param  bool colon
 * @return p
 */
func (p *Item) SetColon(colon bool) *Item {
	p.Colon = colon
	return p
}

/**
 * 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
 *
 * @param  string extra
 * @return p
 */
func (p *Item) SetExtra(extra string) *Item {
	p.Extra = extra
	return p
}

/**
 * 配合 validateStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
 *
 * @param  bool hasFeedback
 * @return p
 */
func (p *Item) SetHasFeedback(hasFeedback bool) *Item {
	p.HasFeedback = hasFeedback
	return p
}

/**
 * 配合 help 属性使用，展示校验状态图标，建议只配合 Input 组件使用
 *
 * @param  bool help
 * @return p
 */
func (p *Item) SetHelp(help string) *Item {
	p.Help = help
	return p
}

/**
 * 为 true 时不带样式，作为纯字段控件使用
 *
 * @return p
 */
func (p *Item) SetNoStyle() *Item {
	p.NoStyle = true
	return p
}

/**
 * label 标签的文本
 *
 * @param  string label
 * @return p
 */
func (p *Item) SetLabel(label string) *Item {
	p.Label = label

	return p
}

/**
 * 标签文本对齐方式
 *
 * @param  left|right align
 * @return p
 */
func (p *Item) SetLabelAlign(align string) *Item {
	p.LabelAlign = align
	return p
}

/**
 * label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。
 * 你可以通过 Form 的 labelCol 进行统一设置。当和 Form 同时设置时，以 Item 为准
 *
 * @param  array|p col
 * @return p
 */
func (p *Item) SetLabelCol(col interface{}) *Item {
	p.LabelCol = col
	return p
}

/**
 * 字段名，支持数组
 *
 * @param  string name
 * @return p
 */
func (p *Item) SetName(name string) *Item {
	p.Name = name
	return p
}

/**
 * 是否必填，如不设置，则会根据校验规则自动生成
 *
 * @return p
 */
func (p *Item) SetRequired() *Item {
	p.Required = true
	return p
}

/**
 * 解析成前端验证规则
 *
 * @param array rules
 * @return array
 */
func parseFrontendRules(rules interface{}, messages interface{}) interface{} {
	result := false

	return result
}

/**
 * 设置前端验证规则
 *
 * @param void
 *
 * @return void
 */
func setFrontendRules() interface{} {
	frontendRules := ""

	return frontendRules
}

/**
 * 校验规则，设置字段的校验逻辑
 *
 * @param  array|p rules
 * @return p
 */
func (p *Item) SetRules(rules interface{}, messages interface{}) *Item {
	p.Rules = rules
	p.RuleMessages = messages

	// 设置前端规则
	setFrontendRules()
	return p
}

/**
 * 校验规则，只在创建表单提交时生效
 *
 * @param  array|p rules
 * @return p
 */
func (p *Item) SetCreationRules(rules interface{}, messages interface{}) *Item {
	p.CreationRules = rules
	p.CreationRuleMessages = messages

	// 设置前端规则
	setFrontendRules()
	return p
}

/**
 * 校验规则，只在更新表单提交时生效
 *
 * @param  array|p rules
 * @return p
 */
func (p *Item) SetUpdateRules(rules interface{}, messages interface{}) *Item {
	p.UpdateRules = rules
	p.UpdateRuleMessages = messages

	// 设置前端规则
	setFrontendRules()
	return p
}

/**
 * 子节点的值的属性，如 Switch 的是 "checked"
 *
 * @param  string valuePropName
 * @return p
 */
func (p *Item) SetValuePropName(valuePropName string) *Item {
	p.ValuePropName = valuePropName
	return p
}

/**
 * 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。
 * 你可以通过 Form 的 wrapperCol 进行统一设置。当和 Form 同时设置时，以 Item 为准。
 *
 * @param  array|p col
 * @return p
 */
func (p *Item) SetWrapperCol(col interface{}) *Item {
	p.WrapperCol = col
	return p
}

/**
 * 设置保存值。
 *
 * @param  array|string
 * @return p
 */
func (p *Item) SetValue(value interface{}) *Item {
	p.Value = value
	return p
}

/**
 * 设置默认值。
 *
 * @param  array|string
 * @return p
 */
func (p *Item) SetDefault(value interface{}) *Item {
	p.DefaultValue = value
	return p
}

/**
 * 是否禁用状态，默认为 false
 *
 * @param  bool disabled
 * @return object
 */
func (p *Item) SetDisabled(disabled bool) *Item {
	p.Disabled = disabled
	return p
}

/**
 * 是否忽略保存到数据库，默认为 false
 *
 * @param  bool ignore
 * @return object
 */
func (p *Item) SetIgnore(ignore bool) *Item {
	p.Ignore = ignore
	return p
}

/**
 * 表单联动
 *
 * @param  mix value
 * @return object
 */
//  func (p *Item) SetWhen(...value interface{}) *Item {
// 	 if(len(value) == 2) {
// 		 operator = "="
// 		 option = value[0]
// 		 whenItem["body"] = value[1]()
// 	 } elseif(len(value) == 3) {
// 		 operator = value[0]
// 		 option = value[1]
// 		 whenItem["body"] = value[2]()
// 	 }

// 	 switch (operator) {
// 		 case "=":
// 			 whenItem["condition"] = "<%=String(" . name . ") === '" . option . "' %>"
// 		   break
// 		 case ">":
// 			 whenItem["condition"] = "<%=String(" . name . ") > '" . option . "' %>"
// 		   break
// 		 case "<":
// 			 whenItem["condition"] = "<%=String(" . name . ") < '" . option . "' %>"
// 		   break
// 		 case "<=":
// 			 whenItem["condition"] = "<%=String(" . name . ") <= '" . option . "' %>"
// 		   break
// 		 case ">=":
// 			 whenItem["condition"] = "<%=String(" . name . ") => '" . option . "' %>"
// 		   break
// 		 case "has":
// 			 whenItem["condition"] = "<%=(String(" . name . ").indexOf('" . option . "') !=-1) %>"
// 		   break
// 		 case "in":
// 			 whenItem["condition"] = "<%=(" . json_encode(option) . ".indexOf(" . name . ") !=-1) %>"
// 		   break
// 		 default:
// 			 whenItem["condition"] = "<%=String(" . name . ") === '" . option . "' %>"
// 		   break
// 	 }

// 	 whenItem["condition_name"] = name
// 	 whenItem["condition_operator"] = operator
// 	 whenItem["condition_option"] = option

// 	 whenItem[] = whenItem
// 	 when["component"] = "when"
// 	 when["items"] = whenItem

// 	 when = when
// 	 return p
//  }

/**
 * Specify that the element should be hidden from the index view.
 *
 * @param  \Closure|bool  callback
 * @return p
 */
func (p *Item) HideFromIndex(callback bool) *Item {
	p.ShowOnIndex = !callback

	return p
}

/**
 * Specify that the element should be hidden from the detail view.
 *
 * @param  \Closure|bool  callback
 * @return p
 */
func (p *Item) HideFromDetail(callback bool) *Item {
	p.ShowOnDetail = !callback

	return p
}

/**
 * Specify that the element should be hidden from the creation view.
 *
 * @param  \Closure|bool  callback
 * @return p
 */
func (p *Item) HideWhenCreating(callback bool) *Item {
	p.ShowOnCreation = !callback

	return p
}

/**
 * Specify that the element should be hidden from the update view.
 *
 * @param  \Closure|bool  callback
 * @return p
 */
func (p *Item) HideWhenUpdating(callback bool) *Item {
	p.ShowOnUpdate = !callback

	return p
}

/**
 * Specify that the element should be hidden from the export file.
 *
 * @param  \Closure|bool  callback
 * @return p
 */
func (p *Item) HideWhenExporting(callback bool) *Item {
	p.ShowOnExport = !callback

	return p
}

/**
 * Specify that the element should be hidden from the import file.
 *
 * @param  \Closure|bool  callback
 * @return p
 */
func (p *Item) HideWhenImporting(callback bool) *Item {
	p.ShowOnImport = !callback

	return p
}

/**
 * Specify that the element should be hidden from the index view.
 *
 * @param  \Closure|bool  callback
 * @return p
 */
func (p *Item) OnIndexShowing(callback bool) *Item {
	p.ShowOnIndex = callback

	return p
}

/**
 * Specify that the element should be hidden from the detail view.
 *
 * @param  \Closure|bool  callback
 * @return p
 */
func (p *Item) OnDetailShowing(callback bool) *Item {
	p.ShowOnDetail = callback

	return p
}

/**
 * Specify that the element should be hidden from the creation view.
 *
 * @param  \Closure|bool  callback
 * @return p
 */
func (p *Item) ShowOnCreating(callback bool) *Item {
	p.ShowOnCreation = callback

	return p
}

/**
 * Specify that the element should be hidden from the update view.
 *
 * @param  \Closure|bool  callback
 * @return p
 */
func (p *Item) ShowOnUpdating(callback bool) *Item {
	p.ShowOnUpdate = callback

	return p
}

/**
 * Specify that the element should be hidden from the export file.
 *
 * @param  \Closure|bool  callback
 * @return p
 */
func (p *Item) ShowOnExporting(callback bool) *Item {
	p.ShowOnExport = callback

	return p
}

/**
 * Specify that the element should be hidden from the import file.
 *
 * @param  \Closure|bool  callback
 * @return p
 */
func (p *Item) ShowOnImporting(callback bool) *Item {
	p.ShowOnImport = callback

	return p
}

/**
 * Specify that the element should only be shown on the index view.
 *
 * @return p
 */
func (p *Item) OnlyOnIndex() *Item {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

/**
 * Specify that the element should only be shown on the detail view.
 *
 * @return p
 */
func (p *Item) OnlyOnDetail() *Item {
	p.ShowOnIndex = false
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

/**
 * Specify that the element should only be shown on forms.
 *
 * @return p
 */
func (p *Item) OnlyOnForms() *Item {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

/**
 * Specify that the element should only be shown on export file.
 *
 * @return p
 */
func (p *Item) OnlyOnExport() *Item {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = false

	return p
}

/**
 * Specify that the element should only be shown on import file.
 *
 * @return p
 */
func (p *Item) OnlyOnImport() *Item {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = true

	return p
}

/**
 * Specify that the element should be hidden from forms.
 *
 * @return p
 */
func (p *Item) ExceptOnForms() *Item {
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = true

	return p
}

/**
 * Check for showing when updating.
 *
 * @return bool
 */
func (p *Item) IsShownOnUpdate() bool {
	return p.ShowOnUpdate
}

/**
 * Check showing on index.
 *
 * @return bool
 */
func (p *Item) IsShownOnIndex() bool {
	return p.ShowOnIndex
}

/**
 * Check showing on detail.
 *
 * @return bool
 */
func (p *Item) IsShownOnDetail() bool {
	return p.ShowOnDetail
}

/**
 * Check for showing when creating.
 *
 * @return bool
 */
func (p *Item) IsShownOnCreation() bool {
	return p.ShowOnCreation
}

/**
 * Check for showing when exporting.
 *
 * @return bool
 */
func (p *Item) IsShownOnExport() bool {
	return p.ShowOnExport
}

/**
 * Check for showing when importing.
 *
 * @return bool
 */
func (p *Item) IsShownOnImport() bool {
	return p.ShowOnImport
}

/**
 * 设置为可编辑列
 *
 * @param  bool  editable
 * @return p
 */
func (p *Item) SetEditable(editable bool) interface{} {
	p.Editable = editable

	return p
}

/**
 * 透传表格列的属性
 *
 * @param  mixed  callback
 * @return void
 */
func (p *Item) SetColumn(column interface{}) *Item {
	p.Column = column

	return p
}
