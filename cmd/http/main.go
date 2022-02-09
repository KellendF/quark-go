package main

import (
	"github.com/quarkcms/quark-go/internal/http"
	"github.com/quarkcms/quark-go/pkg/framework/foundation"
)

func main() {
	kernel := &http.Kernel{}

	foundation.Singleton("httpkernel", kernel)

	httpkernel := foundation.Make("httpkernel")

	httpkernel.(interface{ Run() }).Run()
}
