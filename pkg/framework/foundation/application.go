package foundation

import (
	"fmt"
)

type Application struct {
	Container map[string]T
}

// 简单绑定
func (app *Application) Bind() string {
	return "ss"
}

// 单例绑定
func (app *Application) Singleton(abstract string, concrete T) {
	fmt.Println("ss")
}
