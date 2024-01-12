package utils  // Package declaration, utils package for utility functions

import (
	"encoding/json"  // Importing the encoding/json package for JSON encoding and decoding
	"io/ioutil"      // Importing the ioutil package for reading from and writing to streams
	"net/http"       // Importing the net/http package for HTTP functionality
)

// ParseBody function reads and parses the request body into the provided interface
func ParseBody(r *http.Request, x interface{}) {
	// Reading the entire request body
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		// Attempting to unmarshal (parse) the JSON body into the provided interface
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return  // If there is an error during unmarshalling, return without modifying the provided interface
		}
	}
}
