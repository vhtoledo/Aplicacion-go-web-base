package main

import (
	"fmt"
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
	fmt.Fprintf(w, "<h1>%s</h1> <div>%s</div>", p.Title, p.Body)
}

// Funcion princpal
func main() {
	//p1 := &Page{Title: "TestPage", Body: []byte("Esta es una pagina de muestra")}
	//p1.save()

	//p2, _ := loadPage("TestPage")
	//fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", viewHandler)

	//Crear servidor
	log.Fatal(http.ListenAndServe(":8080", nil))
}