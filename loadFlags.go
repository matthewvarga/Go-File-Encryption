package main

import (
	"errors"
	"flag"
	"fmt"
)

/**
* checks if a specified flag was passed.
* name - name of the flag
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

func loadFlags() (password, path *string, mode Mode, e error) {
	// load flags
	passwordFlag := flag.String("p", "", "encryption password")
	pathFlag := flag.String("path", "", "path to file(s)/directory to encrypt or decrypt")
	encryptModeFlag := flag.Bool("e", false, "Specifies you would like to run encrypt mode")
	decryptModeFlag := flag.Bool("d", false, "Specifies you would like to run decrypt mode")
	flag.Parse()

	var m Mode = EncryptMode

	// verify proper flags are provided
	if !isFlagPassed("p") {
		fmt.Println("Missing Flag: -p.")
		return nil, nil, m, errors.New("Missing Flag: -p")
	}
	if !isFlagPassed("path") {
		fmt.Println("Missing Flag: -path.")
		return nil, nil, m, errors.New("Missing Flag: -path")
	}
	if !*encryptModeFlag && !*decryptModeFlag {
		fmt.Println("Missing Mode Flag: -e or -d.")
		return nil, nil, m, errors.New("Missing Mode Flag: -e or -d")
	}
	if *encryptModeFlag && *decryptModeFlag {
		fmt.Println("Too Many Mode Flags: -e and -d.")
		return nil, nil, m, errors.New("Too Many Mode Flags: -e and -d, Ensure Only 1 is Provided")
	}

	if !*encryptModeFlag {
		m = DecryptMode
	}

	return passwordFlag, pathFlag, m, nil
}
