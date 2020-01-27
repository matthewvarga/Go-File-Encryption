package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// given a file path, checks if it is a file, and if so,
// encrypts or decrypts it, depending on the mode
func processFile(filePath string, m Mode, key [32]byte, wg *sync.WaitGroup) {

	defer wg.Done()

	// check what type the specified path is (file or dir)
	fi, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// if the path is a file, and not directory
	if fi.Mode().IsRegular() {

		// read file contents
		fileContents, err := ioutil.ReadFile(filePath)
		if err != nil {
			//fmt.Println(err)
			//os.Exit(1)
			return
		}

		var newFileContents []byte
		// encode or decode file contents
		if m == EncryptMode {
			newFileContents, err = Encrypt(fileContents, &key)
			if err != nil {
				//fmt.Println(err)
				//os.Exit(1)
				return
			}
		} else if m == DecryptMode {
			newFileContents, err = Decrypt(fileContents, &key)
			if err != nil {
				//fmt.Println(err)
				//os.Exit(1)
				return
			}
		}

		// fmt.Println(filePath)
		// write to file
		err = ioutil.WriteFile(filePath, newFileContents, 0644)

	}
}
