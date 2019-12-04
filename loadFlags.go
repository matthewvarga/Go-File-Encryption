package main

import (
	"errors"
	"flag"
	"fmt"
)

/**
* checks if a specified flag was passed.
* name - name of the flag
* TODO - change to take list of flags
 */
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func loadFlags() (password, path *string, e error) {
	// load flags
	passwordFlag := flag.String("p", "", "encryption password")
	pathFlag := flag.String("path", "", "path to file(s)/directory(s) to encrypt")
	flag.Parse()

	// check if any missing flags
	if !isFlagPassed("p") || !isFlagPassed("path") {
		fmt.Println("One or more required flags are missing. Please try again.")
		return nil, nil, errors.New("Missing requied flag")
	}

	return passwordFlag, pathFlag, nil
}
