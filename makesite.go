package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
)

// Page holds all the information we need to generate a new
// HTML page from a text file on the filesystem.
type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content      string
}

func mdToHtml(fileName string) {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	outputfileHtml := strings.Replace(fileName, ".md", ".html", 1)
	md := markdown.ToHTML(fileContents, nil, nil)
	newFile, err := os.Create(outputfileHtml)
	// create the template
	template := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	template.Execute(newFile, md)
	if err != nil {
		panic(err)
	}

}

func main() {
	filePath := "filePath"
	//create flag for file
	filePtr := flag.String("file", "", "file to read")
	//create flag for directory
	dirPtr := flag.String("dir", "", "directory to read from")
	// create a flag for .md file
	mdFlag := flag.String("md", "test.md", "Render a .md file to HTML")

	if *mdFlag != "test.md" {
		mdToHtml(*mdFlag)
		os.Exit(0)
	}
	flag.Parse()

	if *dirPtr != "" {
		files, err := ioutil.ReadDir(*dirPtr)
		if err != nil {
			fmt.Println(err)
		}
		for _, file := range files {
			// read the file
			fileContents, err := ioutil.ReadFile(file.Name())
			if strings.HasSuffix(file.Name(), ".git") {
				err = nil
			}
			if strings.HasSuffix(file.Name(), ".vscode") {
				err = nil
			}
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}
			// if the file is a text file
			if strings.HasSuffix(file.Name(), ".txt") {
				// replace text file with html file
				filepath := strings.Replace(file.Name(), ".txt", ".html", 1)
				page := Page{
					TextFilePath: filePath,
					TextFileName: *filePtr,
					HTMLPagePath: filepath,
					Content:      string(fileContents),
				}
				// create the template t
				t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
				// Create a new, blank HTML file.
				newFile, err := os.Create(filepath)
				if err != nil {
					panic(err)
				}
				// inject the newly created page into the new htmlfile template
				t.Execute(newFile, page)
			}
			continue
		}
	}
}
