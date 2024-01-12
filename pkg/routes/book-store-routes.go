package routes

import (
	"github.com/Anjasfedo/go-book-store/pkg/controllers"
	"github.com/gorilla/mux"
)

var BookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")

	router.HandleFunc("/book/", controller.GetBooks).Methods("GET")

	router.HandleFunc("/book/{bookId}", controller.GetBookById).Methods("GET")

	router.HandleFunc("/book/{bookId}", controller.UpdateBookById).Methods("PUT")

	router.HandleFunc("/book/{bookId}", controller.DeleteBookById).Methods("DELETE")
}
