package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type person struct {
	ID        int    `json:"id"`
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
	http.HandleFunc("/api/v1/search", search)
	http.HandleFunc("/api/v1/all", friends)
	http.HandleFunc("/api/v1/create", create)
	http.HandleFunc("/api/v1/update", update)
	http.HandleFunc("/api/v1/delete", delete)

	log.Fatal(http.ListenAndServe(":8081", nil))

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func friends(w http.ResponseWriter, r *http.Request) {

	collection := persons
	json.NewEncoder(w).Encode(collection)
}
func create(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/api/v1/create" {
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

		var p *person = &person{}

		fmt.Fprintf(w, "Direccion en memoria de person! = %v\n", &person{})
		fmt.Fprintf(w, "Direccion en memoria de p! = %v\n", &p)

		p.ID = id
		p.Name = name
		p.Age = uint8(age)
		p.Cellphone = cellphone

		addFriend(p)

		fmt.Fprintf(w, "Person created! = %v\n", r.PostForm)

	default:
		fmt.Fprintf(w, "Sorry, the supported methods are GET and POST!")
	}

}

func search(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/api/v1/search" {
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

				fmt.Fprintf(w, "ID = %d\n", v.ID)
				fmt.Fprintf(w, "Name = %s\n", v.Name)
				fmt.Fprintf(w, "Age = %d\n", v.Age)
				fmt.Fprintf(w, "Cellphone = %d\n\n", v.Cellphone)
			}
		}

	default:
		fmt.Fprintf(w, "Sorry, the supported methods are GET and POST!")
	}
}

func update(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/api/v1/update" {
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

			if v.ID == id {

				v.Name = name
				v.Age = uint8(age)
				v.Cellphone = cellphone
			}
		}

		fmt.Fprintf(w, "Updated person!")

	default:
		fmt.Fprintf(w, "Sorry, the supported methods are GET and POST!")
	}
}

func delete(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/api/v1/delete" {
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

			if v.ID != id {

				auxP = append(auxP, v)
			}
		}

		persons = auxP

		fmt.Fprintf(w, "Person removed!")

	default:
		fmt.Fprintf(w, "Sorry, the supported methods are GET and POST!")
	}
}
