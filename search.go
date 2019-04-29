package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func decode(data []byte) (interface{}, error) {
	v, _ := getObject(data)
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

type sliceOp func([]json.RawMessage, string) ([]byte, error)

func sliceIdx(obj []json.RawMessage, index string) ([]byte, error) {
	idx, err := strconv.Atoi(index)
	if err != nil || idx < 0 {
		return nil, fmt.Errorf("invalid index:%s", index)
	}
	if len(obj) < idx {
		return nil, nil
	}
	return obj[idx], nil
}

func sliceMap(obj []json.RawMessage, key string) ([]byte, error) {
	var result []json.RawMessage
	for _, item := range obj {
		val, err := getValue(key, item)
		if err == nil {
			result = append(result, val)
		}
	}
	return sliceSerialize(result), nil
}

func min(x, y int) int {
	return map[bool]int{true: x, false: y}[x < y]
}

func sliceSerialize(obj []json.RawMessage) []byte {
	if len(obj) == 0 {
		return nil
	}
	var result bytes.Buffer
	result.WriteByte('[')
	result.Write(obj[0])
	for _, val := range obj[1:] {
		result.WriteByte(',')
		result.Write(val)
	}
	result.WriteByte(']')
	return result.Bytes()
}

// sliceRange returns a slice of the given array with range operator as key
// it assumes the key is a valid range operator and fallbacks to default values if not
// eg:
// 	`2:4` => obj[2:4]
// 	`:4` => obj[0:4]
// 	`2:` => obj[2:len(obj)]
// 	`2ads:4` => obj[0:4]
// 	`2:dsd4` => obj[2:len(obj)]
func sliceRange(obj []json.RawMessage, key string) ([]byte, error) {
	// get range - assumes it is always given a string with : between
	idxs := strings.Split(key, ":")
	first, _ := strconv.Atoi(idxs[0])
	last, err := strconv.Atoi(idxs[1])
	if err != nil {
		last = len(obj)
	}
	first = min(first, len(obj))
	last = min(last, len(obj))

	if first > last {
		return nil, fmt.Errorf("invalid slice index %d > %d", first, last)
	}
	fmt.Println(first, last)
	return sliceSerialize(obj[first:last]), nil
}

func getSliceOperation(op string) sliceOp {
	_, err := strconv.Atoi(op)
	switch {
	case err == nil:
		return sliceIdx
	case strings.Contains(op, ":"):
		return sliceRange
	default:
		return sliceMap
	}
}

func lookupSlice(key string, obj []json.RawMessage) ([]byte, error) {
	op := getSliceOperation(key)
	return op(obj, key)
}

func getValue(key string, data []byte) ([]byte, error) {
	v, err := decode(data)
	if err != nil {
		return data, err
	}
	switch v := v.(type) {
	case *map[string]json.RawMessage:
		data, _ = (*v)[key]
	case *[]json.RawMessage:
		data, err = lookupSlice(key, *v)
	}
	return data, err
}

func lookup(keys []string, data []byte) ([]byte, error) {
	if len(keys) == 0 {
		return data, nil
	}
	if len(data) == 0 {
		return nil, nil
	}
	data, err := getValue(keys[0], data)
	if err != nil {
		return data, err
	}
	return lookup(keys[1:], data)
}
