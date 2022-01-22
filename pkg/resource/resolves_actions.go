package resource

// 列表行为
func (p *Resource) IndexActions(resource ResourceInterface) interface{} {
	return resource.Actions()
}

// 行为接口
func (p *Resource) Actions() interface{} {
	return nil
}
