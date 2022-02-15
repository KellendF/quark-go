package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/models"
	"github.com/quarkcms/quark-go/pkg/ui/admin/utils"
)

type AdminMiddleware struct{}

// 后台中间件
func (p *AdminMiddleware) Handle(c *fiber.Ctx) error {
	guardName := utils.Admin(c, "guard_name")
	if guardName != "admin" {
		return c.SendStatus(401)
	}

	// 管理员id
	adminId := utils.Admin(c, "id")
	if adminId == nil {
		return c.SendStatus(401)
	}

	if adminId.(float64) != 1 {
		permissions := (&models.Admin{}).GetPermissionsViaRoles(adminId.(float64))
		if permissions == nil {
			return c.SendStatus(403)
		}

		hasPermission := false
		for _, v := range permissions {
			if "/"+v.Name == c.Path() {
				hasPermission = true
			}
		}

		if !hasPermission {
			return c.SendStatus(403)
		}
	}

	return c.Next()
}
