package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

// Page holds all the information we need to generate a new
// HTML page from a text file on the filesystem.
type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content      string
}

func main() {
	filePath := "filePath"
	// read the file
	fileContents, err := ioutil.ReadFile("first-post.txt")
	//create flag for file
	filePtr := flag.String("file", "", "file to read")
	flag.Parse()
	// replace the .txt with .html
	htmlpath := func() string {
		return strings.Replace(*filePtr, ".txt", ".html", -1)
	}
	// creating a new Page object
	page := Page{
		TextFilePath: filePath,
		TextFileName: "first",
		HTMLPagePath: htmlpath(),
		Content:      string(fileContents),
	}
	// create the template t
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	// Create a new, blank HTML file.
	newFile, err := os.Create(htmlpath())
	if err != nil {
		panic(err)
	}
	// inject the newly created page into the new htmlfile named new.html
	t.Execute(newFile, page)
}
