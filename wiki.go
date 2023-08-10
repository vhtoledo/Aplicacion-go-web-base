package main

import (
	"fmt"
	"os"
)

//Estrcutura de la pagina
type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// Funcion para cargar las paginas
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

// Funcion princpal
func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("Esta es una pagina de muestra")}
	p1.save()

	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}