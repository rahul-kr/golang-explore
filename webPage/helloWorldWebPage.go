package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloworld)
	http.ListenAndServe(":8080", nil)
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello There")
	custTemplate, err := template.ParseFiles("helloWorld.html")
	if err != nil {

	}
	err = custTemplate.Execute(w, nil)
	if err != nil {

	}
}
