package main

import (
	"html/template"
	"log"
	"net/http"
)

type Param struct {
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/", helloworld)
	http.ListenAndServe(":8080", nil)
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello There")
	P := Param{Name: "Foo", Age: 20}
	custTemplate, err := template.ParseFiles("helloWorld.html")
	if err != nil {
		log.Println("Error in file parsing")
	}
	err = custTemplate.Execute(w, P)
	if err != nil {

	}
}
