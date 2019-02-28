package main

import (
	"io/ioutil"
	"log"
	"os"
)

// getInput reads input from the input device ,panics if fail
// currently it reads from stdin
func getInput() ([]byte, error) {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Panicf("reading input error : %v", err)
	}
	defer os.Stdin.Close()
	return data, nil
}
