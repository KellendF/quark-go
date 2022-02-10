package config

import (
	"github.com/quarkcms/quark-go/pkg/framework/env"
)

var App = map[string]interface{}{
	// 应用名称
	"name": "QuarkGo v0.0.1",

	// 服务地址
	"host": "127.0.0.1",

	// 服务端口
	"port": "8080",

	// 令牌加密key，默认自动生成，如果设置绝对不可泄漏
	"key": "123456",

	// 服务提供者
	"providers": []interface{}{
		&env.Provider{},
	},
}
