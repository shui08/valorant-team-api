// player-controller.go:
// the controllers package contains all handler functions corresponding to the
// routes defined in pkg/routes.
package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shui08/valorant-team-api/pkg/models"
	"github.com/shui08/valorant-team-api/pkg/utils"
)

const (
	CONTENT_TYPE     = "Content-Type"
	JSON             = "application/json"
	RIOT_ID          = "riotid"
	DELETION_ERR     = "Failed to delete players"
	NO_REQUEST_BODY  = "No request body"
	PLAYER_NOT_FOUND = "Player not found"
)

// this function is a handler for GET requests to the /players endpoint. it
// takes in w, a ResponseWriter, which allows us to directly interact with
// the HTTP response, and it also takes in r, a pointer to a request.
func GetAllPlayers(w http.ResponseWriter, r *http.Request) {

	// this sets the "Content-Type" header of the HTTP response to JSON format.
	w.Header().Set(CONTENT_TYPE, JSON)

	// query for all players in the database and return those players in a
	// slice, which we will define to be `players.`
	players := models.GetAllPlayers()

	// encode players into JSON format and write to w.
	if err := utils.Write(w, players); err != nil {
		http.Error(w, "Error writing response: "+err.Error(), http.StatusInternalServerError)
		log.Printf("GetAllPlayers: error writing response: %v", err)
	}
}

// this function is a handler for GET requests to the /players/{riotid} endpoint
func GetPlayerByID(w http.ResponseWriter, r *http.Request) {

	// this sets the "Content-Type" header of the HTTP response to JSON format.
	w.Header().Set(CONTENT_TYPE, JSON)

	// mux.Vars(r) takes in a Request and returns any route variables for the
	// request as a map. for this specific request, we would extract whatever
	// the client put in for {riotid} in the /players/{riotid} route pattern.
	params := mux.Vars(r)
	riotID := params[RIOT_ID]

	// fetching the player with the specified riotID from the database
	// (see models.GetPlayerByID())
	player, result := models.GetPlayerByID(riotID)
	if result.Error != nil {
		http.Error(w, "Error fetching player: "+result.Error.Error(), http.StatusInternalServerError)
		log.Printf("GetPlayerByID: error fetching player with riotid %s: %v", riotID, result.Error)
		return
	}
	// if player not found, respond with not found error
	if player.RiotID == "" {
		http.Error(w, PLAYER_NOT_FOUND, http.StatusNotFound)
		log.Printf("GetPlayerByID: player with riotid %s not found", riotID)
		return
	}

	// encode player into JSON format and write to w.
	if err := utils.Write(w, player); err != nil {
		http.Error(w, "Error writing response: "+err.Error(), http.StatusInternalServerError)
		log.Printf("GetPlayerByID: error writing response for player %s: %v", riotID, err)
	}
}

// this function is a handler for POST requests to the /players endpoint.
func CreatePlayer(w http.ResponseWriter, r *http.Request) {

	// sets the "Content-Type" header of the HTTP response to JSON format.
	w.Header().Set(CONTENT_TYPE, JSON)

	// creating a pointer to a default player instance (we use a pointer because
	// the AddPlayer instance method takes in a pointer receiver)
	player := &models.Player{}

	// decoding the JSON from the body of the request and storing the data in
	// the player instance
	if err := utils.ParseBody(r, player); err != nil {
		http.Error(w, "Error parsing request body: "+err.Error(), http.StatusBadRequest)
		log.Printf("CreatePlayer: error parsing request body: %v", err)
		return
	}

	// making sure riotid isn't empty
	if player.RiotID == "" {
		http.Error(w, "riotid is required and cannot be empty", http.StatusBadRequest)
		log.Println("CreatePlayer: riotid is missing in the request body")
		return
	}

	// adding the player to the database (see models.AddPlayer())
	if err := player.AddPlayer(); err != nil {
		http.Error(w, "Error adding player: "+err.Error(), http.StatusInternalServerError)
		log.Printf("CreatePlayer: error adding player to database: %v", err)
		return
	}

	// encode player into JSON format and write to w.
	if err := utils.Write(w, player); err != nil {
		http.Error(w, "Error writing response: "+err.Error(), http.StatusInternalServerError)
		log.Printf("CreatePlayer: error writing response: %v", err)
	}
}

// this function is a handler for DELETE requests to the /players endpoint.
func DeleteAllPlayers(w http.ResponseWriter, r *http.Request) {

	// sets the "Content-Type" header of the HTTP response to JSON format.
	w.Header().Set(CONTENT_TYPE, JSON)

	// deleting all players from the database (see models.DeleteAll()) and
	// storing them in a slice
	deletedPlayers, err := models.DeleteAll()

	// if an error occurs while deleting all players, send an error in the
	// response, rather than showing a list of players that weren't actually
	// deleted
	if err != nil {
		http.Error(w, DELETION_ERR+": "+err.Error(), http.StatusInternalServerError)
		log.Printf("DeleteAllPlayers: error deleting players: %v", err)
		return
	}

	// encode deletedPlayers into JSON format and write to w.
	if err := utils.Write(w, deletedPlayers); err != nil {
		http.Error(w, "Error writing response: "+err.Error(), http.StatusInternalServerError)
		log.Printf("DeleteAllPlayers: error writing response: %v", err)
	}
}

// this function is a handler for DELETE requests to the /players/{riotid}
// endpoint.
func DeletePlayer(w http.ResponseWriter, r *http.Request) {

	// sets the "Content-Type" header of the HTTP response to JSON format.
	w.Header().Set(CONTENT_TYPE, JSON)

	// extract riotid from the route pattern
	params := mux.Vars(r)
	riotID := params[RIOT_ID]

	// delete the player with the corresponding RiotID from the database and
	// store that player in deletedPlayer (see models.DeletePlayer)
	deletedPlayer, err := models.DeletePlayer(riotID)
	if err != nil {
		http.Error(w, "Error deleting player: "+err.Error(), http.StatusInternalServerError)
		log.Printf("DeletePlayer: error deleting player with riotid %s: %v", riotID, err)
		return
	}

	// encode deletedPlayer into JSON format and write to w.
	if err := utils.Write(w, deletedPlayer); err != nil {
		http.Error(w, "Error writing response: "+err.Error(), http.StatusInternalServerError)
		log.Printf("DeletePlayer: error writing response: %v", err)
	}
}

// this function is a handler for PUT requests to the /players/{riotid} endpoint
func UpdatePlayer(w http.ResponseWriter, r *http.Request) {

	// sets the "Content-Type" header of the HTTP response to JSON format.
	w.Header().Set(CONTENT_TYPE, JSON)

	// create a new pointer to an empty Player, then parse the request body into
	// that Player struct
	updatedPlayer := &models.Player{}
	err := utils.ParseBody(r, updatedPlayer)

	// if there is an error parsing the request body, then it is likely that the
	// request does not have a body
	if err != nil {
		http.Error(w, NO_REQUEST_BODY+": "+err.Error(), http.StatusBadRequest)
		log.Printf("UpdatePlayer: error parsing request body: %v", err)
		return
	}

	// extract riotid from the route pattern
	params := mux.Vars(r)
	riotID := params[RIOT_ID]

	// fetch the existing player in the database with the corresponding RiotID
	// and store that data in existingPlayer (type Player) and db (type gorm.DB)
	existingPlayer, db := models.GetPlayerByID(riotID)
	if db.Error != nil {
		http.Error(w, "Error fetching player: "+db.Error.Error(), http.StatusInternalServerError)
		log.Printf("UpdatePlayer: error fetching player with riotid %s: %v", riotID, db.Error)
		return
	}
	// if existingPlayer's RiotID field is an empty string, then
	// models.GetPlayerByID could not find a record with the specified RiotID
	// and instead populated existingPlayer with zero values.
	if existingPlayer.RiotID == "" {
		http.Error(w, PLAYER_NOT_FOUND, http.StatusNotFound)
		log.Printf("UpdatePlayer: player with riotid %s not found", riotID)
		return
	}

	// this finds existingPlayer in the database and updates its record to
	// the new values in updatedPlayer. this also ignores zero values in
	// updatePlayer, so that any unspecified fields (which will default to
	// having a zero value) will not override existing fields in existingPlayer.
	updateResult := db.Model(&existingPlayer).Updates(updatedPlayer)
	if updateResult.Error != nil {
		http.Error(w, "Error updating player: "+updateResult.Error.Error(), http.StatusInternalServerError)
		log.Printf("UpdatePlayer: error updating player with riotid %s: %v", riotID, updateResult.Error)
		return
	}

	// encode existingPlayer into JSON format and write to w.
	if err := utils.Write(w, existingPlayer); err != nil {
		http.Error(w, "Error writing response: "+err.Error(), http.StatusInternalServerError)
		log.Printf("UpdatePlayer: error writing response: %v", err)
	}
}
