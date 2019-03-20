package main

import (
	"io"
	"io/ioutil"
	"log"
)

// getInput reads input from the input device ,panics if fail
// currently it reads from stdin
func readStream(inputStream io.Reader) ([]byte, error) {
	data, err := ioutil.ReadAll(inputStream)
	if err != nil {
		log.Panicf("reading input error : %v", err)
	}
	return data, nil
}
