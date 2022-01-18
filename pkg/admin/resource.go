package admin

type Resource struct {
	Title string
}

type IResource interface {
	Fields() interface{}
	Actions() interface{}
}
