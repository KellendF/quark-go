package app

import (
	indexController "github.com/quarkcms/quark-go/app/http/controllers"
)

func Run() {
	indexController.Index()
}
