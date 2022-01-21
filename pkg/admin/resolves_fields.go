package admin

// 列表行为
func (p *Resource) IndexFields(resource ResourceInterface) interface{} {
	return resource.Fields()
}

// 字段接口
func (p *Resource) Fields() interface{} {
	return nil
}
