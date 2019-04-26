package main

import (
	"encoding/json"

	"github.com/TylerBrock/colorjson"
)

func formatJSON(data []byte) ([]byte, error) {
	obj, ok := getObject(data)
	if !ok {
		return data, nil
	}
	if err := json.Unmarshal(data, &obj); err != nil {
		return data, nil
	}
	f := colorjson.NewFormatter()
	f.Indent = 2
	return f.Marshal(obj)
}
