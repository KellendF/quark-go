package main

import (
	"fmt"

	"github.com/quarkcms/quark-go/internal/http"
	"github.com/quarkcms/quark-go/pkg/framework/foundation"
	"github.com/quarkcms/quark-go/pkg/framework/support/facades"
)

func main() {

	fmt.Println(facades.Env.Get("APP_NAME"))

	kernel := &http.Kernel{}

	foundation.Singleton("httpkernel", kernel)

	httpkernel := foundation.Make("httpkernel")

	httpkernel.(interface{ Run() }).Run()
}
