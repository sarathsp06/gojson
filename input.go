package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func getInputStream() io.ReadCloser {
	return os.Stdin
}

// getInput reads input from the input device ,panics if fail
// currently it reads from stdin
func getInput(inputStream io.ReadCloser) ([]byte, error) {
	defer inputStream.Close()
	data, err := ioutil.ReadAll(inputStream)
	if err != nil {
		log.Panicf("reading input error : %v", err)
	}
	return data, nil
}
