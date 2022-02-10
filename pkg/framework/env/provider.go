package env

import (
	"github.com/quarkcms/quark-go/pkg/framework/foundation"
)

type Provider struct{}

// 注册服务
func (p *Provider) Register(key string, value string) {
	env := &Env{}
	foundation.Singleton("env", env)
}
