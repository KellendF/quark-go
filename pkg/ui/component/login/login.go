package login

type Component struct {
	Component   string                   `json:"component"`
	Api         string                   `json:"api"`
	Redirect    string                   `json:"redirect"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	CaptchaUrl  string                   `json:"captchaUrl"`
	Copyright   string                   `json:"copyright"`
	Links       []map[string]interface{} `json:"links"`
}

// 登录接口
func (p *Component) SetApi(api string) *Component {
	p.Api = api
	return p
}

// 登录后跳转地址
func (p *Component) SetRedirect(redirect string) *Component {
	p.Redirect = redirect
	return p
}

// 标题
func (p *Component) SetTitle(title string) *Component {
	p.Title = title
	return p
}

// 描述
func (p *Component) SetDescription(description string) *Component {
	p.Description = description
	return p
}

// 验证码链接
func (p *Component) SetCaptchaUrl(captchaUrl string) *Component {
	p.CaptchaUrl = captchaUrl
	return p
}

// 页脚版权信息
func (p *Component) SetCopyright(copyright string) *Component {
	p.Copyright = copyright
	return p
}

// 页脚友情链接
func (p *Component) SetLinks(links []map[string]interface{}) *Component {
	p.Links = links
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "login"

	return p
}
