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

// getStreamData reads input from the input device ,panics if fail
func getInput() ([]byte, error) {
	inputStream := getInputStream()
	defer inputStream.Close()
	data, err := ioutil.ReadAll(inputStream)
	if err != nil {
		log.Panicf("reading input error : %v", err)
	}
	return data, nil
}
