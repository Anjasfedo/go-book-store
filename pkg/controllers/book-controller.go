package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Anjasfedo/go-book-store/pkg/models"
	"github.com/Anjasfedo/go-book-store/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBooks()

	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "pkglication/json")

	w.WriteHeader(http.StatusOK)

	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookId := params["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails, _ := models.GetBookById(ID)

	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")

	w.WriteHeader(http.StatusOK)

	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}

	utils.ParseBody(r, CreateBook)

	b := CreateBook.CreateBook()

	res, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "pkglication/json")

	w.WriteHeader(http.StatusOK)

	w.Write(res)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookId := params["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	book := models.DeleteBookById(ID)

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")

	w.WriteHeader(http.StatusOK)

	w.Write(res)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}

	utils.ParseBody(r, updateBook)

	params := mux.Vars(r)

	bookId := params["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails, db := models.GetBookById(ID)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")

	w.WriteHeader(http.StatusOK)

	w.Write(res)
}
