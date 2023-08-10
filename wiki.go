package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
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

// Funcion para responder al cliente (visualizar y cargar pagina)
func viewHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	
	renderTemplate(w, "view", p)
}

// Funcion para editar paginas 
func editHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/edit/"):]
	p, _ := loadPage(title)

	renderTemplate(w, "edit", p)
}

// Funcion para renderizar plantillas
func renderTemplate(w http.ResponseWriter, tmpl string, p*Page){
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

// Funcion princpal
func main() {

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)

	//Crear servidor
	log.Fatal(http.ListenAndServe(":8080", nil))
}