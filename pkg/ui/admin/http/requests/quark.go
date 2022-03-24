package requests

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/admin"
	"github.com/quarkcms/quark-go/pkg/ui/admin/utils"
	"gorm.io/gorm"
)

// 初始化
func init() {

	// 自动生成权限信息
	utils.SetPermissions(resourceToPermission())
}

type Quark struct{}

// 资源
func (p *Quark) Resource(c *fiber.Ctx) interface{} {
	var resourceInstance interface{}
	for _, provider := range admin.Providers {
		providerName := reflect.TypeOf(provider).String()

		// 包含字符串
		if find := strings.Contains(providerName, "*resources."); find {
			structName := strings.Replace(providerName, "*resources.", "", -1)

			if strings.ToLower(structName) == strings.ToLower(c.Params("resource")) {
				// 初始化实例
				resourceInstance = provider.(interface{ Init() interface{} }).Init()
			}
		}
	}

	return resourceInstance
}

// 模型
func (p *Quark) NewModel(resourceInstance interface{}) *gorm.DB {
	return resourceInstance.(interface{ NewModel(interface{}) *gorm.DB }).NewModel(resourceInstance)
}

// 资源转换成权限
func resourceToPermission() []string {
	permission := []string{}
	routes := []string{
		"api/admin/dashboard/:dashboard",
		"api/admin/:resource/index",
		"api/admin/:resource/editable",
		"api/admin/:resource/action/:uriKey",
		"api/admin/:resource/create",
		"api/admin/:resource/store",
		"api/admin/:resource/edit",
		"api/admin/:resource/edit/values",
		"api/admin/:resource/save",
	}

	for _, provider := range admin.Providers {
		providerName := reflect.TypeOf(provider).String()
		fmt.Print(providerName)

		// 处理仪表盘
		if find := strings.Contains(providerName, "*dashboard."); find {
			structName := strings.Replace(providerName, "*dashboard.", "", -1)
			for _, v := range routes {
				if strings.Contains(v, ":dashboard") {
					permission = append(permission, strings.Replace(v, ":dashboard", strings.ToLower(structName), -1))
				}
			}
		}

		// 处理资源
		if find := strings.Contains(providerName, "*resources."); find {
			structName := strings.Replace(providerName, "*resources.", "", -1)
			for _, v := range routes {
				if strings.Contains(v, ":resource") {
					permission = append(permission, strings.Replace(v, ":resource", strings.ToLower(structName), -1))

					// 处理行为
					// if strings.Contains(v, ":uriKey") {

					// 	// 初始化实例
					// 	resourceInstance := provider.(interface{ Init() interface{} }).Init()
					// 	actions := resourceInstance.(interface {
					// 		Actions(c *fiber.Ctx) interface{}
					// 	}).Actions(c)

					// 	for _, av := range actions.([]interface{}) {

					// 		// uri唯一标识
					// 		uriKey := av.(interface {
					// 			GetUriKey(interface{}) string
					// 		}).GetUriKey(av)

					// 		permission = append(permission, strings.Replace(v, ":uriKey", uriKey, -1))

					// 	}
					// }
				}
			}
		}
	}

	return []string{}
}
