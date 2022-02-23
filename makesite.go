package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"

	// "path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
)

type Content struct {
	Header     string
	Paragraphs []paragraph
}
type paragraph struct {
	Data string
}

// Read in first-post.txt
func readFile(fileName string) []string {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileData := strings.Split(string(fileContents), "\n")
	return fileData
}

func mdToHtml(fileName string) {
	// read markdown file
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	if strings.HasSuffix(fileName, ".md") {
		// replace markdown extension with html extension
		htmlFileName := strings.Replace(fileName, ".md", ".html", 1)
		md := markdown.ToHTML(fileContent, nil, nil)
		newFile, err := os.Create(htmlFileName)
		if err != nil {
			panic(err)
		}
		// write header to html file
		_, headererr := newFile.WriteString("<!doctype html><html lang='en'><head><meta charset='utf-8'><title>Untitled Custom SSG</title></head><body>")
		if headererr != nil {
			panic(err)
		}
		// write markdown to html file
		_, contenterr := newFile.Write(md)
		if contenterr != nil {
			panic(err)
		}
		// write footer to html file
		_, footererr := newFile.WriteString("</body></html>")
		if footererr != nil {
			panic(err)
		}
		newFile.Close()
	}
}

func convertToHtml(newFileName string) {
	if strings.HasSuffix(newFileName, ".txt") {
		// replace text extension with html extension
		htmlFileName := strings.Replace(newFileName, ".txt", ".html", 1)
		fileToConvert := newFileName
		fileData := readFile(fileToConvert)
		header := fileData[0]
		//Create new Content struct
		var bodyContent []paragraph
		for line := 1; line < len(fileData); line++ {
			if fileData[line] != "" {
				convertedFileData := fileData[line]
				newParagraph := paragraph{Data: convertedFileData}
				bodyContent = append(bodyContent, newParagraph)
			}
		}
		//pass data to template
		htmlTemplate := Content{Header: header, Paragraphs: bodyContent}
		templateParse := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
		newFile, err := os.Create(htmlFileName)
		if err != nil {
			panic(err)
		}
		templateParse.Execute(newFile, htmlTemplate)
	}
}

func main() {
	//dir flag
	dirPtr := flag.String("dir", ".", "Directory to parse")
	//file flag
	filePtr := flag.String("file", "first-post.txt", "file to convert to html")
	//md flag
	mdPtr := flag.String("md", "test.md", "markdown file to convert to html")
	// parse flags
	flag.Parse()

	if *mdPtr != "test.md" {
		mdToHtml(*mdPtr)
		os.Exit(0)
	}
	//output all txt files in current directory
	dir := *dirPtr
	if *dirPtr != "." {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			convertToHtml(file.Name())
		}
		os.Exit(0)
	}
	//generate output file name
	if *filePtr != "first-post.txt" {
		outputFile := *filePtr
		convertToHtml(outputFile)
		os.Exit(0)
	}

	mdToHtml(*mdPtr)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		convertToHtml(file.Name())
	}
	textFiles := *filePtr
	convertToHtml(textFiles)
}
