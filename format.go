package main

import (
	"encoding/json"

	"github.com/TylerBrock/colorjson"
)

func formatJSON(data []byte) ([]byte, error) {
	var obj interface{} = struct{}{}
	if err := json.Unmarshal(data, &obj); err != nil {
		return data, err
	}
	f := colorjson.NewFormatter()
	f.Indent = 2
	return f.Marshal(obj)
}
