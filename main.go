package main

import (
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

const keySeperator = "."

func main() {
	key := getKey()
	data, err := getInput(getInputStream())
	if err != nil {
		fmt.Printf("error reading input: %s", err)
		return
	}
	keys := strings.Split(key, keySeperator)
	if key != "" {
		data, err = lookup(keys, data)
	}
	if err != nil {
		fmt.Printf("error occurred looking up key . Error : %+v ", err)
		return
	}
	formattedJSON, err := formatJSON(data)
	if err != nil {
		fmt.Printf("failed formatting: %s,error: %s", string(data), err)
	}
	fmt.Printf("%s\n", string(formattedJSON))
}
