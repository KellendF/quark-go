package main

import (
	"github.com/quarkcms/quark-go/pkg/framework/foundation"
)

func main() {
	app := &foundation.Application{}

	app.Singleton("httpkernel","love")
}
