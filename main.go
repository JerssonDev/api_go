package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type persona struct {
	Nombre   string `json:"name"`
	Edad     string `json:"age"`
	Telefono string `json:"cel"`
}

var personas []*persona

func addAmigo(p *persona) {
	personas = append(personas, p)
}

func main() {

	http.HandleFunc("/raiz", search)
	http.HandleFunc("/search", search)
	http.HandleFunc("/all", allAmigos)
	http.HandleFunc("/create", create)

	log.Fatal(http.ListenAndServe(":8081", nil))

}

func raiz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvendido!")
	//fmt.Println("Punto Salida: raiz")
}

func allAmigos(w http.ResponseWriter, r *http.Request) {

	collection := &personas

	//fmt.Println("Punto Salida: Todos los articulos")
	json.NewEncoder(w).Encode(collection)
}
func create(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/create" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		fmt.Fprintf(w, "Metodo Post! Create = %v\n", r.PostForm)
		name := r.FormValue("name")
		age := r.FormValue("age")
		cel := r.FormValue("cel")

		p := &persona{}
		p.Nombre = name
		p.Edad = age
		p.Telefono = cel

		addAmigo(p)

	default:
		fmt.Fprintf(w, "Lo siento, los metodos soportados son GET y POST.")
	}

}

func search(w http.ResponseWriter, r *http.Request) {

	//coll := &personas

	//collection := &personas

	if r.URL.Path != "/search" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {

	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		fmt.Fprintf(w, "Metodo Post! Search = %v\n", r.PostForm)
		aux := r.FormValue("name")
		fmt.Fprintf(w, "Name = %s\n", aux)
		/*
			s, _ = json.Marshal(collection)

			for _, v := range p {

				if p.Nombre == aux {
					fmt.Fprintf(w, "Name = %s\n", p.Nombre)
					fmt.Fprintf(w, "Age = %s\n", p.Edad)
					fmt.Fprintf(w, "Cel = %s\n", p.Telefono)

				}

			}*/

	default:
		fmt.Fprintf(w, "Lo siento, los metodos soportados son GET y POST.")
	}
}
