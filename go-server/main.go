package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(response http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(response, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(response, "POST request successful")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(response, "Name = %s\n", name)
	fmt.Fprintf(response, "Address = %s\n", address)
}

func helloHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(response, "404 not found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(response, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(response, "hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8880\n")
	if err := http.ListenAndServe(":8880", nil); err != nil {
		log.Fatal(err)
	}
}
