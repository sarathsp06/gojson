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
		fmt.Printf("Invalid JSON:%s,Error:%s", string(data), err)
	}
	fmt.Printf("%s\n", string(formattedJSON))
}
