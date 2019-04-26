package main

import (
	"encoding/json"
	"strconv"
	"strings"
)

func lookup(key string, data []byte) ([]byte, error) {
	var v, ok = getObject(data)
	keys := strings.Split(key, ".")
	for _, key := range keys {
		if !ok {
			return nil, nil
		}
		switch v := v.(type) {
		case map[string]json.RawMessage:
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			data, ok = v[key]
			if !ok {
				return nil, nil
			}
		case []json.RawMessage:
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			idx, err := strconv.Atoi(key)
			if err != nil {
				return nil, nil
			}
			if len(v) < idx || idx < 0 {
				return nil, nil
			}
			data = v[idx]
		}
		v, ok = getObject(data)
	}
	return data, nil
}
