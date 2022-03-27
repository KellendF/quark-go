package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/framework/token"
)

// 获取管理员Token
func GetAdminToken(c *fiber.Ctx) string {
	header := c.GetReqHeaders()
	getToken := strings.Split(header["Authorization"], " ")

	if len(getToken) != 2 {
		return ""
	}

	return getToken[1]
}

// 获取当前登录管理员信息
func Admin(c *fiber.Ctx, field string) interface{} {
	getToken := GetAdminToken(c)
	userInfo, err := token.Parse(getToken)
	if err != nil {
		return nil
	}

	if value, ok := userInfo[field].(float64); ok {
		return int(value)
	}

	return userInfo[field]
}

// 数据集转换成Tree
func ListToTree(list []map[string]interface{}, pk string, pid string, child string, root int) []interface{} {
	var treeList []interface{}
	for _, v := range list {
		if v[pid] == root {
			childNode := ListToTree(list, pk, pid, child, v[pk].(int))
			if childNode != nil {
				v[child] = childNode
			}
			treeList = append(treeList, v)
		}
	}

	return treeList
}

// Tree转换为有序列表
func TreeToOrderedList(tree []interface{}, level int, field string, child string) []interface{} {
	var list []interface{}
	for _, v := range tree {
		if value, ok := v.(map[string]interface{}); ok {
			value[field] = strings.Repeat("—", level) + value[field].(string)
			list = append(list, value)
			if value[child] != nil {
				if childValue, ok := value[child].([]interface{}); ok {
					children := TreeToOrderedList(childValue, level+1, field, child)
					list = append(list, children...)
				}
			}
		}
	}

	return list
}

// struct转map
func StructToMap(v any) any {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}
	var mapResult any
	json.Unmarshal(jsonBytes, &mapResult)

	return mapResult
}

// 存储权限
var Permissions []string

// 设置权限
func SetPermissions(permissions []string) {
	Permissions = permissions
}

// 获取权限
func GetPermissions() []string {
	return Permissions
}
