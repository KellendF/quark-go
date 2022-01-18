package admin

// 列表行为
func (p *Resource) IndexFields(child IResource) interface{} {
	return child.Fields()
}

// 字段接口
func (p *Resource) Fields() interface{} {
	return nil
}
