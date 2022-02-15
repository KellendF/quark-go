package metrics

type TotalAdmin struct{}

// 初始化
func (p *TotalAdmin) Init() {

}

// 计算数值
func (dashboard *TotalAdmin) Calculate() interface{} {

	return "测试内容"
}
