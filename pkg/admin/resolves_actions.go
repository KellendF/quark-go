package admin

// 列表行为
func (p *Resource) IndexActions(child IResource) interface{} {
	return child.Actions()
}

// 行为接口
func (p *Resource) Actions() interface{} {
	return nil
}
