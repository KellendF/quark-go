package utils

import (
	"encoding/json"
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

// 数据集转换成Tree
func ListToTree(list []interface{}, pk string, pid string, child string, root float64) []interface{} {
	var treeList []interface{}
	for _, v := range list {
		if v.(map[string]interface{})["pid"] == root {
			childNode := ListToTree(list, pk, pid, child, v.(map[string]interface{})[pk].(float64))
			if childNode != nil {
				v.(map[string]interface{})[child] = childNode
			}
			treeList = append(treeList, v)
		}
	}

	return treeList
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
