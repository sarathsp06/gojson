package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pkg/errors"

	"strings"
)

//getObject returns the object into which the json can be decoded
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
		return 0, false

	}
}

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
				return nil, errors.New("invalid key")
			}
		case []json.RawMessage:
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			idx, err := strconv.Atoi(key)
			if err != nil {
				return nil, errors.New("invalid key:" + key)
			}
			if len(v) < idx || idx < 0 {
				return nil, errors.New("invalid index:" + key)
			}
			data = v[idx]
		}
		v, ok = getObject(data)
	}
	return data, nil
}

func getKey() string {
	if len(os.Args) == 1 {
		return ""
	}
	key := strings.TrimSpace(os.Args[1])
	return strings.TrimFunc(key, func(c rune) bool { return c == '.' })
}

func main() {
	key := getKey()
	data, err := getInput()
	if err != nil {
		log.Printf("Error reading input: %s", err)
		return
	}
	if key != "" {
		data, err = lookup(key, data)
		if err != nil {
			log.Println("Error occurred looking up key . Error : ", err)
			return
		}
	}
	formattedJSON, err := formatJSON(data)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s\n", string(formattedJSON))
}
