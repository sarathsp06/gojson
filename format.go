package main

import (
	"encoding/json"

	"github.com/TylerBrock/colorjson"
)

func formatJSON(data []byte) ([]byte, error) {
	f := colorjson.NewFormatter()
	f.Indent = 4
	v, _ := getObject(data)
	if err := json.Unmarshal(data, &v); err != nil {
		return data, err
	}
	formattedJSON, err := f.Marshal(v)
	if err != nil {
		return data, err
	}
	return formattedJSON, nil
}
