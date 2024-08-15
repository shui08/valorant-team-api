// the utils package will hold any necessary helper functions.
package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	DECODING_ERR = "Decoding error occurred"
	ENCODING_ERR = "Encoding error occurred"
)

// this function is for parsing the JSON body of a request into a Go struct. it
// takes in a r, a request, and x, which can be of any type (as specified by the
// usage of interface{}). it also returns an error.
func ParseBody(r *http.Request, x interface{}) {
	// unmarshal the  data from the request's body and store it in the value
	// pointed to by `x`. Decoder.Decode also returns an error, and if an error
	// actually does occur, we will simply print the error.
	err := json.NewDecoder(r.Body).Decode(&x)
	if err != nil {
		fmt.Println(DECODING_ERR)
	}
}

// this function writes a struct, x, to w in JSON format.
func Write(w http.ResponseWriter, x interface{}) {

	// encode the data from x and write it to w. if an error occurs, simply
	// print the error.
	err := json.NewEncoder(w).Encode(x)
	if err != nil {
		fmt.Println(ENCODING_ERR)
	}
}
