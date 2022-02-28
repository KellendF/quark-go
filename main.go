package main

import (
	"embed"

	"github.com/quarkcms/quark-go/internal/http"
)

//go:embed assets/*
var assets embed.FS

func main() {

	// 服务实例
	kernel := &http.Kernel{}

	// 启动服务
	kernel.Run(assets)
}
