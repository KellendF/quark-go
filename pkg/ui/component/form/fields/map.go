package fields

import "github.com/quarkcms/quark-go/pkg/ui/component"

type Map struct {
	Item
	Zoom                int    `json:"zoom"`
	MapKey              string `json:"mapKey"`
	ButtonPosition      string `json:"buttonPosition"`
	AlwaysShowItemLabel bool   `json:"alwaysShowItemLabel"`
}

// 初始化
func (p *Map) Init() *Map {
	p.Component = "mapField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.Value = map[string]string{
		"longitude": "116.397724",
		"latitude":  "39.903755",
	}
	p.Zoom = 14
	p.MapKey = "788e08def03f95c670944fe2c78fa76f"
	p.Style["height"] = 500
	p.Style["width"] = "100%"
	p.Style["marginTop"] = "10px"

	return p
}

// zoom
func (p *Map) SetZoom(zoom int) *Map {
	p.Zoom = zoom
	return p
}

// 高德地图key
func (p *Map) SetMapKey(key string) *Map {
	p.MapKey = key
	return p
}

// 地图宽度
func (p *Map) SetWidth(width interface{}) *Map {
	p.Style["width"] = width
	return p
}

// 地图高度
func (p *Map) SetHeight(height interface{}) *Map {
	p.Style["height"] = height
	return p
}

// 坐标位置
func (p *Map) SetPosition(longitude string, latitude string) *Map {
	p.Value = map[string]string{
		"longitude": longitude,
		"latitude":  latitude,
	}
	return p
}
