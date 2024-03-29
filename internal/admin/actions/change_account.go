package actions

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/framework/hash"
	"github.com/quarkcms/quark-go/pkg/framework/msg"
	"github.com/quarkcms/quark-go/pkg/ui/admin/actions"
	"github.com/quarkcms/quark-go/pkg/ui/admin/utils"
	"gorm.io/gorm"
)

type ChangeAccount struct {
	actions.Action
}

// 执行行为句柄
func (p *ChangeAccount) Handle(c *fiber.Ctx, model *gorm.DB) error {
	data := map[string]interface{}{}
	json.Unmarshal(c.Body(), &data)

	if data["avatar"] != "" {
		data["avatar"], _ = json.Marshal(data["avatar"])
	} else {
		data["avatar"] = nil
	}

	// 加密密码
	if data["password"] != nil {
		data["password"] = hash.Make(data["password"].(string))
	}

	result := model.Where("id", utils.Admin(c, "id")).Updates(data).Error

	if result != nil {
		return msg.Error(c, result.Error(), "")
	}

	return msg.Success(c, "操作成功！", "", "")
}
