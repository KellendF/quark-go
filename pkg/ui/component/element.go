package component

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/go-basic/uuid"
)

type Element struct {
	Key       string                 `json:"key"`
	Component string                 `json:"component"`
	Style     map[string]interface{} `json:"style"`
}

var DEFAULT_KEY = ""
var DEFAULT_KEY_CRYPT = false

// Set the key for the component.
func (p *Element) SetKey(key string, crypt bool) *Element {
	if key == "" {
		p.Key = uuid.New()
	}

	if crypt {
		h := md5.New()
		h.Write([]byte(key))
		p.Key = hex.EncodeToString(h.Sum(nil))
	}

	return p
}

// Set the component.
func (p *Element) SetComponent(component string) *Element {
	p.Component = component

	return p
}

// Set style.
func (p *Element) SetStyle(style map[string]interface{}) *Element {
	p.Style = style

	return p
}
