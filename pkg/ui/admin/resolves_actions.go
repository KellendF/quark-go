package admin

// 列表行为
func (p *Resource) IndexActions(resource interface{}) interface{} {
	return resource.(interface{ Actions() interface{} }).Actions()
}
