package dashboard

import "github.com/gofiber/fiber/v2"

// 资源结构体
type Dashboard struct {
	Title    string
	SubTitle string
}

// 资源接口
type DashboardInterface interface {
	Init()
	HandleInit(dashboard DashboardInterface)
	SetTitle(title string)
	GetTitle() string
	Cards() interface{}
	GetCards(dashboard DashboardInterface) interface{}
	Render(c *fiber.Ctx, dashboard DashboardInterface, content interface{}) interface{}
	DashboardComponentRender(c *fiber.Ctx, dashboard DashboardInterface) interface{}
}

// 初始化资源
func (p *Dashboard) Init() {}

// 执行资源初始化
func (p *Dashboard) HandleInit(dashboard DashboardInterface) {
	dashboard.Init()
}

// 设置标题
func (p *Dashboard) SetTitle(title string) {
	p.Title = title
}

// 获取标题
func (p *Dashboard) GetTitle() string {
	return p.Title
}

// 设置子标题
func (p *Dashboard) SetSubTitle(subTitle string) {
	p.SubTitle = subTitle
}

// 获取子标题
func (p *Dashboard) GetSubTitle() string {
	return p.SubTitle
}

// 卡片列表
func (p *Dashboard) Cards() interface{} {
	return nil
}

// 获取卡片列表
func (p *Dashboard) GetCards(dashboard DashboardInterface) interface{} {

	// $cards = $this->cards();

	// $colNum = 0;
	// $rows = $cols = [];

	// foreach ($cards as $key => $card) {
	// 	$colNum = $colNum + $card->col;

	// 	$cardItem = Card::body(
	// 		$card->calculate($request)
	// 	);

	// 	$cols[] = Col::span($card->col)->body($cardItem);
	// 	if($colNum%24 === 0) {
	// 		$row = Row::gutter(8);
	// 		if($key !== 1) {
	// 			$row = $row->style(['marginTop' => '20px']);
	// 		}
	// 		$rows[] = $row->body($cols);
	// 		$cols = [];
	// 	}
	// }

	// if($cols) {
	// 	$row = Row::gutter(8);
	// 	if($colNum > 24) {
	// 		$row = $row->style(['marginTop' => '20px']);
	// 	}
	// 	$rows[] = $row->body($cols);
	// }

	// return $rows;

	return dashboard.Cards()
}
