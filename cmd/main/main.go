package main  // Package declaration, main package is the entry point for executable programs

import (
	"fmt"     // Importing the fmt package for formatted I/O
	"log"     // Importing the log package for logging
	"net/http" // Importing the net/http package for HTTP server functionality

	"github.com/Anjasfedo/go-book-store/pkg/routes" // Importing custom routes package
	"github.com/gorilla/mux"                        // Importing the Gorilla Mux router package
	_ "github.com/jinzhu/gorm/dialects/mysql"       // Importing the MySQL dialect for GORM, but not directly used in the code
)

func main() {
	r := mux.NewRouter() // Creating a new Gorilla Mux router

	routes.BookStoreRoutes(r) // Registering routes using custom function from imported routes package

	http.Handle("/", r) // Handling root path with the created router

	fmt.Printf("Starting Server on Port 8000\n") // Printing a message indicating the server is starting

	log.Fatal(http.ListenAndServe(":8000", r)) // Starting the HTTP server on port 8000, log.Fatal logs any errors and terminates the program if an error occurs
}
