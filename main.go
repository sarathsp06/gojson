package main

import (
	"fmt"
	"log"
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
	data, err := getInput()
	if err != nil {
		log.Printf("Error reading input: %s", err)
		return
	}
	keys := strings.Split(key, keySeperator)
	data, err = lookup(keys, data)
	if err != nil {
		log.Printf("Error occurred looking up key . Error : %+v ", err)
		return
	}

	formattedJSON, err := formatJSON(data)
	if err != nil {
		log.Printf("Invalid JSON: %s,Error: %s", string(data), err)
	}
	fmt.Printf("%s : %s\n", key, string(formattedJSON))
}
