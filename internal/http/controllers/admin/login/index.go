package login

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/http/controllers/tool/captcha"
	"github.com/quarkcms/quark-go/pkg/db"
	"github.com/quarkcms/quark-go/pkg/hash"
	"github.com/quarkcms/quark-go/pkg/token"
)

// 请求结构体
type Request struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Captcha  string `json:"captcha" form:"captcha"`
}

// admins表结构体
type Admin struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

// 表结构体
type Component struct {
	Component   string              `json:"component"`
	Api         string              `json:"api"`
	Redirect    string              `json:"redirect"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	CaptchaUrl  string              `json:"captchaUrl"`
	Copyright   string              `json:"copyright"`
	Links       []map[string]string `json:"links"`
}

// 登录页面
func Show(c *fiber.Ctx) error {

	Component := &Component{
		Component:   "login",
		Api:         "admin/login",
		Redirect:    "/index?api=admin/dashboard/index",
		Title:       "QuarkAdmin",
		Description: "信息丰富的世界里，唯一稀缺的就是人类的注意力",
		CaptchaUrl:  "api/admin/captcha",
		Copyright:   "版权所有",
		Links: []map[string]string{
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
		},
	}

	return c.JSON(Component)
}

// 登录方法
func Login(c *fiber.Ctx) error {
	request := new(Request)

	if err := c.BodyParser(request); err != nil {
		return err
	}

	if !captcha.Check(request.Captcha) {
		return c.JSON(fiber.Map{
			"status": "error",
			"msg":    "验证码错误！",
		})
	}

	if request.Username == "" || request.Password == "" {
		return c.JSON(fiber.Map{
			"status": "error",
			"msg":    "用户名或密码不能为空！",
		})
	}

	admin := Admin{}
	db.Conn().Where("username = ?", request.Username).First(&admin)

	dict := make(map[string]interface{})
	dict["id"] = admin.Id
	dict["avatar"] = admin.Avatar
	dict["nickname"] = admin.Nickname
	dict["username"] = admin.Username
	dict["guard_name"] = "admin"

	getToken, _ := token.Make(dict)

	if !hash.Check(admin.Password, request.Password) {
		return c.JSON(fiber.Map{
			"status": "error",
			"msg":    "用户名或密码错误！",
		})
	} else {
		return c.JSON(fiber.Map{
			"status": "success",
			"msg":    "登录成功！",
			"data": map[string]string{
				"id":       strconv.Itoa(admin.Id),
				"avatar":   admin.Avatar,
				"nickname": admin.Nickname,
				"username": admin.Username,
				"token":    getToken,
			},
		})
	}
}
