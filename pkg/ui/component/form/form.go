package form

import "github.com/quarkcms/quark-go/pkg/ui/component"

type Component struct {
	component.Element
	Title              string                 `json:"title"`
	Width              string                 `json:"width"`
	Colon              bool                   `json:"colon"`
	Values             map[string]interface{} `json:"values"`
	InitialValues      map[string]interface{} `json:"initialValues"`
	LabelAlign         string                 `json:"labelAlign"`
	Name               string                 `json:"name"`
	Preserve           bool                   `json:"preserve"`
	RequiredMark       bool                   `json:"requiredMark"`
	ScrollToFirstError bool                   `json:"scrollToFirstError"`
	Size               string                 `json:"size"`
	DateFormatter      string                 `json:"dateFormatter"`
	Layout             string                 `json:"layout"`
	LabelCol           map[string]interface{} `json:"labelCol"`
	WrapperCol         map[string]interface{} `json:"wrapperCol"`
	ButtonWrapperCol   map[string]interface{} `json:"buttonWrapperCol"`
	Api                string                 `json:"api"`
	ApiType            string                 `json:"apiType"`
	TargetBlank        bool                   `json:"targetBlank"`
	InitApi            string                 `json:"initApi"`
	Body               interface{}            `json:"body"`
	Actions            []interface{}          `json:"actions"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "form"
	p.Colon = true
	p.LabelAlign = "right"
	p.Preserve = true
	p.RequiredMark = true
	p.Size = "default"
	p.DateFormatter = "string"
	p.Layout = "horizontal"
	p.LabelCol = map[string]interface{}{
		"span": 4,
	}
	p.WrapperCol = map[string]interface{}{
		"span": 20,
	}
	p.ButtonWrapperCol = map[string]interface{}{
		"offset": 4,
		"span":   20,
	}
	p.ApiType = "POST"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 配置表单标题
func (p *Component) SetTitle(title string) *Component {
	p.Title = title
	return p
}

// 配置表单宽度
func (p *Component) SetWidth(width string) *Component {
	p.Width = width

	return p
}

//  配置 Form.Item 的 colon 的默认值。表示是否显示 label 后面的冒号 (只有在属性 layout 为 horizontal 时有效)
func (p *Component) SetColon(colon bool) *Component {
	p.Colon = colon

	return p
}

//  解析initialValue
func (p *Component) SetParseInitialValue(item interface{}, initialValues map[string]interface{}) map[string]interface{} {
	//  if(isset($item->name)) {

	// 	 if(isset($item->defaultValue)) {
	// 		 $value = $item->defaultValue;
	// 	 }

	// 	 if(isset($initialValues[$item->name])) {
	// 		 $value = $initialValues[$item->name];
	// 	 }

	// 	 if(isset($item->value)) {
	// 		 $value = $item->value;
	// 	 }
	//  }

	//  return $value ?? null;
	return map[string]interface{}{}
}

// 表单默认值，只有初始化以及重置时生效
func (p *Component) SetInitialValues(initialValues map[string]interface{}) *Component {
	//  $data = $initialValues;

	//  foreach ($this->items as $item) {
	// 	 $value = $this->parseInitialValue($item,$initialValues);

	// 	 if($value !== null) {
	// 		 $data[$item->name] = $value;
	// 	 }
	//  }

	//  foreach ($data as $key => $value) {
	// 	 if(is_string($value)) {
	// 		 if(count(explode('[',$value))>1 || count(explode('{',$value))>1) {
	// 			 $value = json_decode($value, true);
	// 		 }
	// 	 }

	// 	 $data[$key] = $value;
	//  }

	//  $this->initialValues = $data;

	return p
}

//  label 标签的文本对齐方式,left | right
func (p *Component) SetLabelAlign(labelAlign string) *Component {
	p.LabelAlign = labelAlign

	return p
}

//  表单名称，会作为表单字段 id 前缀使用
func (p *Component) SetName(name string) *Component {
	p.Name = name

	return p
}

//  当字段被删除时保留字段值
func (p *Component) SetPreserve(preserve bool) *Component {
	p.Preserve = preserve

	return p
}

//  必选样式，可以切换为必选或者可选展示样式。此为 Form 配置，Form.Item 无法单独配置
func (p *Component) SetRequiredMark(requiredMark bool) *Component {
	p.RequiredMark = requiredMark

	return p
}

//  提交失败自动滚动到第一个错误字段
func (p *Component) SetScrollToFirstError(scrollToFirstError bool) *Component {
	p.ScrollToFirstError = scrollToFirstError

	return p
}

//  设置字段组件的尺寸（仅限 antd 组件）,small | middle | large
func (p *Component) SetSize(size string) *Component {
	p.Size = size

	return p
}

//  自动格式数据，例如 moment 的表单,支持 string 和 number 两种模式
func (p *Component) SetDateFormatter(dateFormatter string) *Component {
	p.DateFormatter = dateFormatter

	return p
}

// 表单布局，horizontal|vertical
func (p *Component) SetLayout(layout string) *Component {

	if layout == "vertical" {
		p.LabelCol = nil
		p.WrapperCol = nil
		p.ButtonWrapperCol = nil
	}

	p.Layout = layout

	return p
}

// label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}
func (p *Component) SetLabelCol(labelCol map[string]interface{}) *Component {
	p.LabelCol = labelCol

	return p
}

// 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol
func (p *Component) SetWrapperCol(wrapperCol map[string]interface{}) *Component {
	if p.Layout == "vertical" {
		panic("If layout set vertical mode,can't set wrapperCol!")
	}

	p.WrapperCol = wrapperCol

	return p
}

//  表单按钮布局样式,默认：['offset' => 2, 'span' => 22 ]
func (p *Component) SetButtonWrapperCol(buttonWrapperCol map[string]interface{}) *Component {
	if p.Layout == "vertical" {
		panic("If layout set vertical mode,can't set buttonWrapperCol!")
	}

	p.ButtonWrapperCol = buttonWrapperCol

	return p
}

// 表单提交的接口链接
func (p *Component) SetApi(api string) *Component {
	p.Api = api

	return p
}

//  表单提交接口的类型
func (p *Component) SetApiType(apiType string) *Component {
	p.ApiType = apiType

	return p
}

//  提交表单的数据是否打开新页面，只有在GET类型的时候有效
func (p *Component) SetTargetBlank(targetBlank bool) *Component {
	p.TargetBlank = targetBlank

	return p
}

//  获取表单数据
func (p *Component) SetInitApi(initApi string) *Component {
	p.InitApi = initApi

	return p
}

//  表单项
func (p *Component) SetBody(items interface{}) *Component {
	p.Body = items

	return p
}

/**
 * 解析保存提交数据库前的值
 *
 * @param array $rules
 * @return array
 */
func (p *Component) parseSubmitData(data map[string]interface{}) interface{} {
	items := p.Body

	//  foreach ($items as $key => $item) {

	// 	 // 删除忽略的值
	// 	 if($item->ignore) {
	// 		 if(isset($data[$item->name])) {
	// 			 unset($data[$item->name]);
	// 		 }
	// 	 }
	//  }

	//  foreach ($data as $key => $value) {
	// 	 if(is_array($value)) {
	// 		 $data[$key] = json_encode($value);
	// 	 }
	//  }

	//  return $data;

	return items
}

/**
 * 行为
 *
 * @param array $actions
 * @return $this
 */
func (p *Component) SetActions(actions []interface{}) *Component {
	p.Actions = actions

	return p
}
