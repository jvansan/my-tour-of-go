package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// Page described how html data will be stored
type Page struct {
	Title string
	// Use []byte because this is how io/ioutils expects it
	Body []byte
}

// Persistent storage of page data
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) //0600 is r+w permission
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	// Pass error on
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	// Create a page and save it
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()

	// Load the saved page
	p2, err := loadPage("TestPage")
	if err != nil {
		log.Fatal("Could not load page ", err)
	}
	fmt.Println(string(p2.Body))
}
