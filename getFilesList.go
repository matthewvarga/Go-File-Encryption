package main

import (
	"os"
	"path/filepath"
)

// returns a list of files in the given path.
// includes all subdirectory files if they exist
// ex: given the following file structure:
// testDir
//     test1.txt
//     testDir2
//          test2.txt
// getFilesList("./testDir") =
//    ["./testDir, ./testDir/test1.txt", "./testDir/testDir2/",
//     "./testDir/testDir2/test2.txt"]
func getFilesList(path string) []string {
	var files []string
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		files = append(files, p)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
