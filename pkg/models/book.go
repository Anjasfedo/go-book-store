package models  // Package declaration, models package for defining data models

import (
	"github.com/Anjasfedo/go-book-store/pkg/config"  // Importing the custom config package for database connection
	"github.com/jinzhu/gorm"  // Importing GORM, a popular Object-Relational Mapping (ORM) library for Go
)

var db *gorm.DB  // Variable to hold the database connection pool

// Book struct defines the data model for books with additional metadata provided by gorm.Model
type Book struct {
	gorm.Model   // Embedding gorm.Model for metadata fields like ID, CreatedAt, UpdatedAt, and DeletedAt
	Name        string `gorm:""json:"name"`  // Book name with JSON tag for serialization
	Author      string `json:"author"`        // Book author with JSON tag for serialization
	Publication string `json:"publication"`   // Book publication with JSON tag for serialization
}

func init() {
	config.Connect()  // Initializing the database connection by calling the Connect function from the config package

	db = config.GetDB()  // Getting the database connection pool from the config package

	db.AutoMigrate(&Book{})  // Auto-migrating the Book model to create the corresponding database table
}

// CreateBook method creates a new book record in the database
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)  // Creating a new record in the database for the book

	db.Create(&b)  // Creating the book record in the database

	return b  // Returning the created book
}

// GetBooks function retrieves all books from the database
func GetBooks() []Book {
	var Books []Book  // Slice to hold the retrieved books

	db.Find(&Books)  // Querying the database to retrieve all books

	return Books  // Returning the list of books
}

// GetBookById function retrieves a book by its ID from the database
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book  // Variable to hold the retrieved book

	db := db.Where("ID=?", Id).Find(&getBook)  // Querying the database to retrieve the book by ID

	return &getBook, db  // Returning the retrieved book and the database connection
}

// DeleteBookById function deletes a book by its ID from the database
func DeleteBookById(Id int64) Book {
	var book Book  // Variable to hold the book to be deleted

	db.Where("ID=?", Id).Delete(book)  // Deleting the book from the database by ID

	return book  // Returning the deleted book
}
