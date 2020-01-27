package main

import (
	"crypto/sha256"
	"os"
	"sync"
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

	var wg sync.WaitGroup

	// process the file(s)
	for _, file := range files {
		wg.Add(1)

		go processFile(file, mode, hash, &wg)
	}

	wg.Wait()
}
