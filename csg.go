package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	files, err := ioutil.ReadDir(".")
	check(err)

	for _, file := range files {
		// Iterate through all non-hidden directories
		if file.IsDir() && (strings.Index(file.Name(), ".") != 0) {
			// Create a separate reference file for each directory at this level
			createRefFile(file)
		}
	}
}

func createRefFile(dir fs.FileInfo) {
	// Create the file
	file, err := os.Create(dir.Name() + ".md")
	check(err)

	defer file.Close()

	// Add the top-level heading
	file.WriteString("# " + dir.Name() + "\n\n")
}
