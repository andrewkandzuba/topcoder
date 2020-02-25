package wiki

import (
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save(root string) error {
	filename := root + "/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func LoadPage(root string, title string) *Page {
	filename := root + "/" + title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}
}