// the utils package will hold any necessary helper functions.
package utils

import (
	"encoding/json"
	"net/http"
)

// this function is for parsing the JSON body of a request into a Go struct. it
// takes in a r, a request, and x, which can be of any type (as specified by the
// usage of interface{}). it also returns an error.
func ParseBody(r *http.Request, x interface{}) error {
	// unmarshal the  data from the request's body and store it in the value
	// pointed to by `x`. Decoder.Decode also returns an error, and if an error
	// actually does occur, we will return the error and let the caller handle
	// it.
	err := json.NewDecoder(r.Body).Decode(&x)
	if err != nil {
		return err
	}
	return nil
}

// this function writes a struct, x, to w in JSON format.
func Write(w http.ResponseWriter, x interface{}) error {

	// encode the data from x and write it to w. if an error occurs, return the
	// error and let the caller handle it.
	err := json.NewEncoder(w).Encode(x)
	if err != nil {
		return err
	}
	return nil
}
