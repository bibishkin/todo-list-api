package main

import (
	"log"
	"os"
)

var infoLog = log.New(logFile("info.txt"), "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

var errorLog = log.New(logFile("errors.txt"), "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

func logFile(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	return file
}
