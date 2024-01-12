package controller  // Package declaration, controller package for handling HTTP requests and responses

import (
	"encoding/json"   // Importing the encoding/json package for JSON encoding and decoding
	"fmt"             // Importing the fmt package for formatted I/O
	"github.com/Anjasfedo/go-book-store/pkg/models"  // Importing the custom models package
	"github.com/Anjasfedo/go-book-store/pkg/utils"   // Importing the custom utils package
	"github.com/gorilla/mux"   // Importing the Gorilla Mux router package
	"net/http"  // Importing the net/http package for HTTP functionality
	"strconv"   // Importing the strconv package for string conversions
)

var NewBook models.Book  // Declaring a global variable of type Book from the models package

// GetBooks function handles the HTTP GET request to retrieve all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBooks()  // Calling the GetBooks function from the models package to get all books

	res, _ := json.Marshal(newBooks)  // Marshaling the books into JSON format

	w.Header().Set("Content-Type", "application/json")  // Setting the response content type to JSON

	w.WriteHeader(http.StatusOK)  // Setting the HTTP status code to 200 (OK)

	w.Write(res)  // Writing the JSON response to the HTTP response writer
}

// GetBookById function handles the HTTP GET request to retrieve a book by its ID
func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)  // Getting the URL parameters using Gorilla Mux

	bookId := params["bookId"]  // Extracting the book ID from the parameters

	ID, err := strconv.ParseInt(bookId, 0, 0)  // Converting the book ID to an integer
	if err != nil {
		fmt.Println("Error while parsing")  // Printing an error message if there is an issue parsing the book ID
	}

	bookDetails, _ := models.GetBookById(ID)  // Calling the GetBookById function from the models package to get book details by ID

	res, _ := json.Marshal(bookDetails)  // Marshaling the book details into JSON format

	w.Header().Set("Content-Type", "application/json")  // Setting the response content type to JSON

	w.WriteHeader(http.StatusOK)  // Setting the HTTP status code to 200 (OK)

	w.Write(res)  // Writing the JSON response to the HTTP response writer
}

// CreateBook function handles the HTTP POST request to create a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}  // Creating a new Book instance

	utils.ParseBody(r, CreateBook)  // Parsing the request body and populating the Book instance

	b := CreateBook.CreateBook()  // Creating the book using the CreateBook method from the models package

	res, _ := json.Marshal(b)  // Marshaling the created book into JSON format

	w.Header().Set("Content-Type", "application/json")  // Setting the response content type to JSON

	w.WriteHeader(http.StatusOK)  // Setting the HTTP status code to 200 (OK)

	w.Write(res)  // Writing the JSON response to the HTTP response writer
}

// DeleteBookById function handles the HTTP DELETE request to delete a book by its ID
func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)  // Getting the URL parameters using Gorilla Mux

	bookId := params["bookId"]  // Extracting the book ID from the parameters

	ID, err := strconv.ParseInt(bookId, 0, 0)  // Converting the book ID to an integer
	if err != nil {
		fmt.Println("Error while parsing")  // Printing an error message if there is an issue parsing the book ID
	}

	book := models.DeleteBookById(ID)  // Calling the DeleteBookById function from the models package to delete a book by ID

	res, _ := json.Marshal(book)  // Marshaling the deleted book into JSON format

	w.Header().Set("Content-Type", "application/json")  // Setting the response content type to JSON

	w.WriteHeader(http.StatusOK)  // Setting the HTTP status code to 200 (OK)

	w.Write(res)  // Writing the JSON response to the HTTP response writer
}

// UpdateBookById function handles the HTTP PUT request to update a book by its ID
func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}  // Creating a new Book instance for the updated book details

	utils.ParseBody(r, updateBook)  // Parsing the request body and populating the Book instance with updated details

	params := mux.Vars(r)  // Getting the URL parameters using Gorilla Mux

	bookId := params["bookId"]  // Extracting the book ID from the parameters

	ID, err := strconv.ParseInt(bookId, 0, 0)  // Converting the book ID to an integer
	if err != nil {
		fmt.Println("Error while parsing")  // Printing an error message if there is an issue parsing the book ID
	}

	bookDetails, db := models.GetBookById(ID)  // Getting the existing book details by ID from the models package

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name  // Updating the book name if provided in the request
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author  // Updating the book author if provided in the request
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication  // Updating the book publication if provided in the request
	}

	db.Save(&bookDetails)  // Saving the updated book details to the database

	res, _ := json.Marshal(bookDetails)  // Marshaling the updated book details into JSON format

	w.Header().Set("Content-Type", "application/json")  // Setting the response content type to JSON

	w.WriteHeader(http.StatusOK)  // Setting the HTTP status code to 200 (OK)

	w.Write(res)  // Writing the JSON response to the HTTP response writer
}
