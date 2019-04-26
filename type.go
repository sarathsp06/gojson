package main

import (
	"encoding/json"
)

// getObject returns the object into which the json can be decoded
func getObject(jsn []byte) (interface{}, bool) {
	if len(jsn) == 0 {
		return json.RawMessage{}, false
	}
	switch jsn[0] {
	case '{':
		return make(map[string]json.RawMessage), true
	case '[':
		return []json.RawMessage{}, true
	case 'n':
		return json.RawMessage{}, false
	case '"', '\'':
		return "", false
	default:
		return 0.0, false
	}
}
