package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func decode(data []byte) (interface{}, error) {
	v, _ := getObject(data)
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

func lookupSlice(key string, obj []json.RawMessage) ([]byte, error) {
	idx, err := strconv.Atoi(key)
	if err != nil || idx < 0 {
		return nil, fmt.Errorf("invalid index:%s", key)
	}
	if len(obj) < idx {
		return nil, nil
	}
	return obj[idx], nil
}

func lookup(key []string, data []byte) ([]byte, error) {
	fmt.Println(key, string(data))
	if len(key) == 0 {
		return data, nil
	}
	if len(data) == 0 {
		return nil, nil
	}
	v, err := decode(data)
	if err != nil {
		return data, err
	}

	switch v := v.(type) {
	case *map[string]json.RawMessage:
		data, _ = (*v)[key[0]]
	case *[]json.RawMessage:
		data, err = lookupSlice(key[0], *v)
	}
	if err != nil {
		return data, err
	}
	return lookup(key[1:], data)
}
