package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mayowa1305/bookStore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
