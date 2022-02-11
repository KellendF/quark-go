package utils

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/framework/token"
)

// 获取当前登录管理员信息
func Admin(c *fiber.Ctx, field string) interface{} {

	// token认证
	header := c.GetReqHeaders()
	getToken := strings.Split(header["Authorization"], " ")

	if len(getToken) != 2 {
		return nil
	}

	userInfo, err := token.Parse(getToken[1])
	if err != nil {
		fmt.Println(err)
	}

	return userInfo[field]
}
