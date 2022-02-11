package main

import (
	"github.com/quarkcms/quark-go/internal/http"
)

func main() {

	// 服务实例
	kernel := &http.Kernel{}

	// 启动服务
	kernel.Run()
}
