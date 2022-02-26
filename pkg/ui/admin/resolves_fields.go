package admin

// 列表行为
func (p *Resource) IndexFields(resource interface{}) interface{} {
	return resource.(interface{ Fields() interface{} }).Fields()
}
