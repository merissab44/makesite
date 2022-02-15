package main

import (
        "html/template"
        "os"
		"io/ioutil"
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
	// creating a new Page object 
    page := Page{
      TextFilePath: filePath,
      TextFileName: "first",
      HTMLPagePath: "first-post.html",
      Content:      string(fileContents),
    }
   // create the template t
    t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
    // Create a new, blank HTML file.
    newFile, err := os.Create("first-post.html")
    if err != nil {
          panic(err)
    }
    // inject the newly created page into the new htmlfile named new.html
    t.Execute(newFile, page)
}