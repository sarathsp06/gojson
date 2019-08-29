package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"strings"
)

func getKey() string {
	if len(os.Args) == 1 {
		return ""
	}
	key := strings.TrimSpace(os.Args[1])
	return strings.TrimFunc(key, func(c rune) bool { return c == '.' })
}

func splitKey(key string, seperator rune) ([]string, error) {
	r := csv.NewReader(strings.NewReader(key))
	r.Comma = seperator
	keys, err := r.Read()
	if err != nil {
		return nil, err
	}
	return keys, err
}

const keySeperator = '.'

func main() {
	key := getKey()
	data, err := getInput(getInputStream())
	if err != nil {
		fmt.Printf("error reading input: %s", err)
		return
	}
	keys, err := splitKey(key, keySeperator)
	if err != nil {
		fmt.Printf("invalid key. Error : %+v ", err)
		return
	}
	if key != "" {
		data, err = lookup(keys, data)
	}
	if err != nil {
		fmt.Printf("error occurred looking up key. Error : %+v ", err)
		return
	}
	formattedJSON, err := formatJSON(data)
	if err != nil {
		fmt.Printf("failed formatting: %s,error: %s", string(data), err)
	}
	fmt.Printf("%s\n", string(formattedJSON))
}
