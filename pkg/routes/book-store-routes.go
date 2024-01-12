package routes

import (
	"fmt"

	"github.com/Anjasfedo/go-book-store/pkg/controllers"
	"github.com/gorilla/mux"
)

var BookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controller.createBook).Methods("POST")

	router.HandleFunc("/book/", controller.getBooks).Methods("GET")

	router.HandleFunc("/book/{bookId}", controller.getBookById).Methods("GET")

	router.HandleFunc("/book/{bookId}", controller.updateBookById).Methods("PUT")

	router.HandleFunc("/book/{bookId}", controller.deleteBookById).Methods("DELETE")
}
