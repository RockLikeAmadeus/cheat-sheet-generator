package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var refFileTitle = ""

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

	// set the title for this reference file
	refFileTitle = dir.Name()
	// Add the top-level heading
	refFile.WriteString(generateHeading(1, refFileTitle) + "\n\n")

	// Write the contents for the top-level directory
	writeContents(dir, refFile, ".", 2)
}

/* writeIndex writes the table of contents for the directory dir to the file refFile */
func writeIndex(dir fs.FileInfo, refFile *os.File, path string, level int) {
	// Write the heading
	refFile.WriteString(generateHeading(level, dir.Name()) + "\n\n")

	// Iterate through sub-directories and files
	files, err := ioutil.ReadDir(path + "/" + dir.Name())
	check(err)

	for _, file := range files {
		name := file.Name()
		// If this is a markdown file, chop off the extension
		name = strings.TrimSuffix(name, ".md")
		// Write an internal link for each sub-directory
		refFile.WriteString("[[#" + name + "]]\n")
	}
	refFile.WriteString("\n\n" + generateBackLink() + "\n\n<br><br><br><br><br>\n---\n\n")
	// refFile.WriteString("\n---\n\n")
}

/* writeContents writes the content for the directory dir to the file refFile */
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
			// Write contents for directories
			writeContents(file, refFile, path+"/"+dir.Name(), level+1)
		} else if strings.HasSuffix(file.Name(), ".md") { // Only write the contents of markdown files
			// Write contents for markdown files
			headingTitle := strings.TrimSuffix(file.Name(), ".md")
			refFile.WriteString(generateHeading(6, headingTitle) + "\n\n")
			content, err := ioutil.ReadFile(path + "/" + dir.Name() + "/" + file.Name())
			check(err)
			refFile.WriteString(string(content) + "\n\n")
			refFile.WriteString(generateBackLink() + "\n\n<br><br><br><br><br>\n---\n\n")
		}
	}
}

func generateHeading(level int, name string) string {
	// Obsidian Markdown doesn't parse h7 or larger
	if level > 6 {
		level = 6
	}
	return strings.Repeat("#", level) + " " + name
}

func generateBackLink() string {
	return "_[[#" + refFileTitle + "|Back to Top]]_"
}
