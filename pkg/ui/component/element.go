package component

type Element struct {
	Key       string                 `json:"key"`
	Component string                 `json:"component"`
	Style     map[string]interface{} `json:"style"`
}
