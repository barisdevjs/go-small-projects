package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found Hello", http.StatusNotFound)
		return
	}

	if !methodChecker(r.Method, "GET") {
		http.Error(w, "Method is not Supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello from Hello Server")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 Not Found Form", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm err: %v", err)
		return
	}

	fmt.Print(r)
	fmt.Fprintf(w, "POST req is successful\n")
	name := r.Form.Get("name")
	address := r.Form.Get("address")
	fmt.Fprintf(w, "Name = %s\n Address = %s", name, address)

}

func methodChecker(method string, comparer string) bool {
	return method == comparer
}

func main() {
	// checkout static file directory
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server listening on 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
