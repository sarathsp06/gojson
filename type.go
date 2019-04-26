package main

import (
	"encoding/json"
)

// getObject returns the object into which the json can be decoded
func getObject(jsn []byte) (interface{}, bool) {
	obj := make(map[string]json.RawMessage)
	if len(jsn) == 0 {
		return &json.RawMessage{}, false
	}
	switch jsn[0] {
	case '{':
		return &obj, true
	case '[':
		return &[]json.RawMessage{}, true
	case 'n':
		return &json.RawMessage{}, false
	case '"', '\'':
		return new(string), false
	default:
		return new(float64), false
	}
}
