package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sid/go-store/pkg/routes"
)

func main() {
    r := mux.NewRouter()
	routes.RegisterbookstoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9001", r))
	if err := http.ListenAndServe("localhost:9001", r); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}