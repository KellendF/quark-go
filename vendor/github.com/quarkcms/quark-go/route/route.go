package route

import (
	"fmt"
	indexController "github.com/quarkcms/quark-go/app/http/controllers"
)

func Get(url, controller) {
	fmt.Println(controller)
}
