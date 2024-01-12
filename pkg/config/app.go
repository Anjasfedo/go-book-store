package config  // Package declaration, config package for managing configurations

import (
	"github.com/jinzhu/gorm"          // Importing GORM, a popular Object-Relational Mapping (ORM) library for Go
	_ "github.com/jinzhu/gorm/dialects/mysql" // Importing MySQL dialect for GORM, the underscore (_) indicates that the imported package is only used for its side effects
)

var (
	db *gorm.DB  // Variable to hold the database connection pool
)

// Connect function establishes a connection to the MySQL database
func Connect() {
	d, err := gorm.Open("mysql", "root:@/bookstore?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)  // Panic if there is an error connecting to the database
	}

	db = d  // Assigning the database connection to the global variable
}

// GetDB function returns the database connection pool
func GetDB() *gorm.DB {
	return db  // Returning the database connection pool
}
