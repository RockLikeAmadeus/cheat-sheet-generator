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
	refFile, err := os.Create(dir.Name() + ".md")
	check(err)

	defer refFile.Close()

	// Add the top-level heading
	refFile.WriteString("# " + dir.Name() + "\n\n")

	// Iterate through sub-directories
	files, err := ioutil.ReadDir(dir.Name())
	check(err)

	// Write the table of contents for each sub-directory
	for _, file := range files {
		if file.IsDir() && (strings.Index(file.Name(), ".") != 0) {
			writeIndex(file, refFile, dir.Name(), 2)
		}
	}
	refFile.WriteString("---\n\n")

	// Write the contents for each sub-directory and file
	for _, file := range files {
		if file.IsDir() && (strings.Index(file.Name(), ".") != 0) {
			writeContents(file, refFile, dir.Name(), 3)
		}
	}
}

/*
writeIndex writes the table of contents for the directory dir to the file refFile
*/
func writeIndex(dir fs.FileInfo, refFile *os.File, path string, level int) {
	heading := strings.Repeat("#", level) + " " + dir.Name() + "\n\n"
	refFile.WriteString(heading)

	// Iterate through sub-directories and files
	files, err := ioutil.ReadDir(path + "/" + dir.Name())
	check(err)

	for _, file := range files {
		// Write an internal link for each sub-directory
		refFile.WriteString("[[#" + file.Name() + "]]\n")
	}
	refFile.WriteString("\n---\n\n")
}

func writeContents(dir fs.FileInfo, refFile *os.File, path string, level int) {
	// Iterate through sub-directories and files
	files, err := ioutil.ReadDir(path + "/" + dir.Name())
	check(err)

	// Write the table of contents for each sub-directory
	for _, file := range files {
		if file.IsDir() && (strings.Index(file.Name(), ".") != 0) {
			writeIndex(file, refFile, path+"/"+dir.Name(), level)
		}
	}
	refFile.WriteString("\n")

	// Write the contents for each sub-directory
	for _, file := range files {
		if file.IsDir() && (strings.Index(file.Name(), ".") != 0) {
			writeContents(file, refFile, path+"/"+dir.Name(), level+1)
		} else {
			refFile.WriteString("Contents Here\n\n")
		}
	}
}
