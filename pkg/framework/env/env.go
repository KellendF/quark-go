package env

import (
	"fmt"

	"github.com/spf13/viper"
)

type Env struct{}

// 设置值
func (p *Env) Set(key string, value string) bool {

	return false
}

// 获取值
func (p *Env) Get(key string) interface{} {
	viper.SetConfigFile(".env") // 指定配置文件路径

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	return viper.Get(key)
}
