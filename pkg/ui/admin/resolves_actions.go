package admin

// 列表行为
func (p *Resource) IndexActions(resourceInstance interface{}) interface{} {
	return resourceInstance.(interface{ Actions() interface{} }).Actions()
}
