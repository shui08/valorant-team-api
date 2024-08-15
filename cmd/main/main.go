// this will be the entry point for the application.
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shui08/valorant-team-api/pkg/routes"
)

const (
	STARTING_SERVER = "Starting server at port 8080"
)

func main() {
	// creating a new Router instance
	r := mux.NewRouter()

	// initializing the routes defined in pkg/routes to r (see
	// routes.InitializeTeamRoutes())
	routes.InitializeTeamRoutes(r)

	// setting r to handle all requests to the root path, which effectively
	// sets r to handle any request made to the server
	http.Handle("/", r)

	// starting the server and telling it to listen on port 8080 while using r
	// to handle any requests. if a non-nil error is returned by ListenAndServe,
	// we will log it and exit the program.
	fmt.Println(STARTING_SERVER)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
