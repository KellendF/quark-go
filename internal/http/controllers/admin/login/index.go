package login

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/http/controllers/tool/captcha"
	"github.com/quarkcms/quark-go/internal/model"
	"github.com/quarkcms/quark-go/pkg/component/login"
	"github.com/quarkcms/quark-go/pkg/hash"
	"github.com/quarkcms/quark-go/pkg/msg"
	"github.com/quarkcms/quark-go/pkg/token"
)

// 请求结构体
type Request struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Captcha  string `json:"captcha" form:"captcha"`
}

// 登录页面
func Show(c *fiber.Ctx) error {

	loginComponent := &login.Component{}

	component := loginComponent.
		SetApi("admin/login").
		SetRedirect("/index?api=admin/dashboard/index").
		SetTitle("QuarkGo").
		SetDescription("信息丰富的世界里，唯一稀缺的就是人类的注意力").
		SetCaptchaUrl("api/admin/captcha").
		SetCopyright("版权所有").
		SetLinks([]map[string]string{
			{
				"title": "迁安信息港",
				"href":  "http://www.qa114.com/",
			},
			{
				"title": "迁安人才网",
				"href":  "http://www.qarc.cn",
			},
			{
				"title": "深蓝科技",
				"href":  "https://www.qasl.cn",
			},
		}).
		JsonSerialize()

	return c.JSON(component)
}

// 登录方法
func Login(c *fiber.Ctx) error {
	request := new(Request)

	if err := c.BodyParser(request); err != nil {
		return err
	}

	if !captcha.Check(request.Captcha) {
		return msg.Error("验证码错误", msg.DEFAULT_URL)
	}

	if request.Username == "" || request.Password == "" {
		return msg.Error("用户名或密码不能为空", msg.DEFAULT_URL)
	}

	model := &model.AdminModel{}
	admin := model.FindByUsername(request.Username)

	// 检验账号和密码
	if !hash.Check(admin.Password, request.Password) {
		return msg.Error("用户名或密码错误", msg.DEFAULT_URL)
	}

	data := make(map[string]interface{})
	data["id"] = admin.Id
	data["avatar"] = admin.Avatar
	data["nickname"] = admin.Nickname
	data["username"] = admin.Username
	data["guard_name"] = "admin"

	// 创建token
	getToken, _ := token.Make(data)
	data["token"] = getToken

	return msg.Success("登录成功", msg.DEFAULT_URL, data)
}
