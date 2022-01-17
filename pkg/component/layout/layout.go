package layout

type Component struct {
	Component     string              `json:"component"`
	Cache         bool                `json:"cache"`
	Title         string              `json:"title"`
	Logo          string              `json:"logo"`
	Loading       bool                `json:"loading"`
	ContentStyle  map[string]string   `json:"contentStyle"`
	HeaderActions []map[string]string `json:"headerActions"`
	Layout        string              `json:"layout"`
	HeaderTheme   string              `json:"headerTheme"`
	SplitMenus    bool                `json:"splitMenus"`
	ContentWidth  string              `json:"contentWidth"`
	NavTheme      string              `json:"navTheme"`
	PrimaryColor  string              `json:"primaryColor"`
	FixedHeader   bool                `json:"fixedHeader"`
	FixSiderbar   bool                `json:"fixSiderbar"`
	IconfontUrl   bool                `json:"iconfontUrl"`
	Locale        string              `json:"locale"`
	SiderWidth    int                 `json:"siderWidth"`
	Collapsed     bool                `json:"collapsed"`
	Menu          interface{}         `json:"menu"`
	Footer        interface{}         `json:"footer"`
	Body          interface{}         `json:"body"`
}

// 是否缓存layout
func (p *Component) SetCache(cache bool) *Component {
	p.Cache = cache
	return p
}

// layout 的左上角 的 title
func (p *Component) SetTitle(title string) *Component {
	p.Title = title
	return p
}

// layout 的左上角的 logo
func (p *Component) SetLogo(logo string) *Component {
	p.Logo = logo
	return p
}

// layout 的加载态
func (p *Component) SetLoading(loading bool) *Component {
	p.Loading = loading
	return p
}

// layout 的内容区 style
func (p *Component) SetContentStyle(contentStyle map[string]string) *Component {
	p.ContentStyle = contentStyle
	return p
}

// layout 的头部行为
func (p *Component) SetHeaderActions(headerActions []map[string]string) *Component {
	p.HeaderActions = headerActions
	return p
}

// layout 的布局模式，side：右侧导航，top：顶部导航，mix：混合模式
func (p *Component) SetLayout(layout string) *Component {
	p.Layout = layout
	return p
}

// layout为mix时，顶部主题 dark | light
func (p *Component) SetHeaderTheme(headerTheme string) *Component {
	p.HeaderTheme = headerTheme
	return p
}

// layout 的内容模式,Fluid：定宽 1200px，Fixed：自适应
func (p *Component) SetContentWidth(contentWidth string) *Component {
	p.ContentWidth = contentWidth
	return p
}

// 导航的主题，'light' | 'dark'
func (p *Component) SetNavTheme(navTheme string) *Component {
	p.NavTheme = navTheme
	return p
}

// 容器控件里面的内容
func (p *Component) SetPrimaryColor(primaryColor string) *Component {
	p.PrimaryColor = primaryColor
	return p
}

// 是否固定导航
func (p *Component) SetFixedHeader(fixedHeader bool) *Component {
	p.FixedHeader = fixedHeader
	return p
}

// 是否固定导航
func (p *Component) SetFixSiderbar(fixSiderbar bool) *Component {
	p.FixSiderbar = fixSiderbar
	return p
}

// todo

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "layout"

	return p
}
