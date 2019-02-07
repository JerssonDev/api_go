package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type person struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Age       uint8  `json:"age"`
	Cellphone int    `json:"cellphone"`
}

var persons []*person

func addFriend(p *person) {
	persons = append(persons, p)
}

func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/search", search)
	http.HandleFunc("/all", friends)
	http.HandleFunc("/create", create)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)

	log.Fatal(http.ListenAndServe(":8081", nil))

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
	//fmt.Println("Punto Salida: raiz")
}

func friends(w http.ResponseWriter, r *http.Request) {

	collection := persons

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

		id, _ := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("name")
		age, _ := strconv.Atoi(r.FormValue("age"))
		cellphone, _ := strconv.Atoi(r.FormValue("cellphone"))

		p := &person{}
		p.Id = id
		p.Name = name
		p.Age = uint8(age)
		p.Cellphone = cellphone

		addFriend(p)

		fmt.Fprintf(w, "Person created! = %v\n", r.PostForm)

	default:
		fmt.Fprintf(w, "Lo siento, los metodos soportados son GET y POST.")
	}

}

func search(w http.ResponseWriter, r *http.Request) {

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

		fmt.Fprintf(w, "Metodo Post! Search = %v\n\n", r.PostForm)
		aux := r.FormValue("name")
		collection := persons

		for _, v := range collection {

			if v.Name == aux {

				fmt.Fprintf(w, "Id = %d\n", v.Id)
				fmt.Fprintf(w, "Name = %s\n", v.Name)
				fmt.Fprintf(w, "Age = %d\n", v.Age)
				fmt.Fprintf(w, "Cellphone = %d\n\n", v.Cellphone)
			}
		}

	default:
		fmt.Fprintf(w, "Lo siento, los metodos soportados son GET y POST.")
	}
}

func update(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/update" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {

	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		fmt.Fprintf(w, "Metodo Post! Update = %v\n", r.PostForm)
		id, _ := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("name")
		age, _ := strconv.Atoi(r.FormValue("age"))
		cellphone, _ := strconv.Atoi(r.FormValue("cellphone"))

		collection := persons

		for _, v := range collection {

			if v.Id == id {

				v.Name = name
				v.Age = uint8(age)
				v.Cellphone = cellphone
			}
		}

		fmt.Fprintf(w, "Updated person!")

	default:
		fmt.Fprintf(w, "Lo siento, los metodos soportados son GET y POST.")
	}
}

func delete(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/delete" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {

	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		fmt.Fprintf(w, "Metodo Post! Update = %v\n", r.PostForm)
		id, _ := strconv.Atoi(r.FormValue("id"))

		collection := persons

		var auxP []*person

		for _, v := range collection {

			if v.Id != id {

				auxP = append(auxP, v)
			}
		}

		persons = auxP

		fmt.Fprintf(w, "Person removed!")

	default:
		fmt.Fprintf(w, "Lo siento, los metodos soportados son GET y POST.")
	}
}
