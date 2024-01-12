package main

import (
	"github.com/Anjasfedo/go-book-store/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	routes.BookStoreRoutes(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8000", r))
}
