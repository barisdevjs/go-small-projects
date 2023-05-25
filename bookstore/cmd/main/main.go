package main

import (
	"bookstore/pkg/config"
	"bookstore/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	config.Connect() // Establish the PostgreSQL database connection

	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r)) // Start the HTTP server on port 8080
	fmt.Print("STARTED")

}
