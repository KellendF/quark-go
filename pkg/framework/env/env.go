package env

import (
	"fmt"

	"github.com/spf13/viper"
)

// 设置值
func Set(key string, value interface{}) {
	viper.SetConfigFile(".env")
	viper.Set(key, value)
}

// 获取值
func Get(key ...string) interface{} {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if len(key) == 2 {
		if viper.Get(key[0]) == nil {
			return key[1]
		}
	}

	return viper.Get(key[0])
}
