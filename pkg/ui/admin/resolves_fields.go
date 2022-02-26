package admin

// 列表行为
func (p *Resource) IndexFields(resourceInstance interface{}) interface{} {
	return resourceInstance.(interface{ Fields() interface{} }).Fields()
}
