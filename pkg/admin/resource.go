package admin

type Resource struct {
	Title string
}

type IResource interface {
	Fields() interface{}
	Actions() interface{}
	Title() string
}

// 列表行为
func (p *Resource) GetTitle(child IResource) string {
	return child.Title()
}
