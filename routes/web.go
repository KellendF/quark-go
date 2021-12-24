package webRoute

import (
	route "github.com/quarkcms/quark-go/route"
)

func Routes() {
	route.Get("/", "IndexController@Index")
}
