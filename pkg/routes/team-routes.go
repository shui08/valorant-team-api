// the routes package contains all the necessary logic for routing requests to
// their corresponding handler functions, which will be defined in the
// controllers package.
package routes

import (
	"github.com/gorilla/mux"
	"github.com/shui08/valorant-team-api/pkg/controllers"
)

// this function is intended to be used in the main.go file, where it will set
// up routes for an existing Router instance. routes will match HTTP requests
// (specified by Route.Methods()) at a certain endpoint (specified in the first
// argument of HandleFunc()) to the correct handler function (specified in the
// second argument of HandleFunc).
func InitializeTeamRoutes(router *mux.Router) {
	// for example, if a HTTP "POST" request hits the /player endpoint, the
	// controllers.CreatePlayer function will be called.
	router.HandleFunc("/player", controllers.CreatePlayer).Methods("POST")
	router.HandleFunc("/player", controllers.GetAllPlayers).Methods("GET")
	router.HandleFunc("/player/{user}", controllers.GetPlayerByUser).Methods("GET")
	router.HandleFunc("/player/{user}", controllers.UpdatePlayer).Methods("PUT")
	router.HandleFunc("/player/{user}", controllers.DeletePlayer).Methods("DELETE")
}
