package main

import (
	//"fmt"
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
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

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// Funcion validar y obtener titulo
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Titulo de página inválido")
	}
	
	return m[2], nil
}

// Funcion para responder al cliente (visualizar y cargar pagina)
func viewHandler(w http.ResponseWriter, r *http.Request){
	title, err := getTitle(w ,r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// Funcion para editar paginas 
func editHandler(w http.ResponseWriter, r *http.Request){
	title, err := getTitle(w ,r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

// Funcion para guardar paginas
func saveHandler(w http.ResponseWriter, r *http.Request){
	title, err := getTitle(w ,r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// Funcion para renderizar plantillas
func renderTemplate(w http.ResponseWriter, tmpl string, p*Page){
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Funcion princpal
func main() {

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	//Crear servidor
	log.Fatal(http.ListenAndServe(":8080", nil))
}