package main

import (
	"crypto/sha256"
	"os"
)

func main() {

	// load flags
	password, path, mode, err := loadFlags()
	if err != nil {
		os.Exit(1)
	}

	// hash password
	hash := sha256.Sum256([]byte(*password))

	// retrieve all files in the provided path
	files := getFilesList(*path)

	// process the file(s)
	for _, file := range files {
		processFile(file, mode, hash)
	}
}
