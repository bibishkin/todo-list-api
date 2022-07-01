package main

import (
	"io"
	"os"
)

func GetHTML(path string) ([]byte, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	html, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return html, nil
}

func GetBody(body io.Reader) ([]byte, error) {
	readBody, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	return readBody, nil
}
