package main

import (
	"encoding/json"

	"github.com/TylerBrock/colorjson"
)

func formatJSON(data []byte) ([]byte, error) {
	var obj map[string]interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		//return as such
		return data, nil
	}
	f := colorjson.NewFormatter()
	f.Indent = 2
	return f.Marshal(obj)
}
