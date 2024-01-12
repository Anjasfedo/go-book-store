package routes  // Package declaration, routes package for defining API routes

import (
	"github.com/Anjasfedo/go-book-store/pkg/controllers"  // Importing the custom controllers package for handling HTTP requests
	"github.com/gorilla/mux"  // Importing the Gorilla Mux router package for defining routes
)

// BookStoreRoutes is a function that defines and configures the routes for the book store API
var BookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")  // Creating a route for creating a new book, using the CreateBook function from the controllers package, with the HTTP method POST

	router.HandleFunc("/book/", controller.GetBooks).Methods("GET")  // Creating a route for retrieving all books, using the GetBooks function from the controllers package, with the HTTP method GET

	router.HandleFunc("/book/{bookId}", controller.GetBookById).Methods("GET")  // Creating a route for retrieving a book by ID, using the GetBookById function from the controllers package, with the HTTP method GET

	router.HandleFunc("/book/{bookId}", controller.UpdateBookById).Methods("PUT")  // Creating a route for updating a book by ID, using the UpdateBookById function from the controllers package, with the HTTP method PUT

	router.HandleFunc("/book/{bookId}", controller.DeleteBookById).Methods("DELETE")  // Creating a route for deleting a book by ID, using the DeleteBookById function from the controllers package, with the HTTP method DELETE
}
