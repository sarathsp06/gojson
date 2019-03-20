package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func selectInputStream() io.ReadCloser {
	return os.Stdin
}

// getStreamData reads input from the input device ,panics if fail
func getStreamData(inputStream io.ReadCloser) ([]byte, error) {
	data, err := ioutil.ReadAll(inputStream)
	if err != nil {
		log.Panicf("reading input error : %v", err)
	}
	return data, nil
}

func getInput() ([]byte, error) {
	inputStream := selectInputStream()
	defer inputStream.Close()
	return getStreamData(inputStream)
}
