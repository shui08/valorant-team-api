// the controllers package contains all handler functions corresponding to the
// routes defined in pkg/routes.
package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shui08/valorant-team-api/pkg/models"
	"github.com/shui08/valorant-team-api/pkg/utils"
)

const (
	CONTENT_TYPE = "Content-Type"
	JSON         = "application/json"
	RIOT_ID      = "riotid"
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
	utils.Write(w, players)
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
	player, _ := models.GetPlayerByID(riotID)

	// encode player into JSON format and write to w.
	utils.Write(w, player)
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
	utils.ParseBody(r, player)

	// adding the player to the database (see models.AddPlayer())
	player.AddPlayer()

	// encode player into JSON format and write to w.
	utils.Write(w, player)
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
	deletedPlayer := models.DeletePlayer(riotID)

	// encode deletedPlayer into JSON format and write to w.
	utils.Write(w, deletedPlayer)
}

// this function is a handler for PUT requests to the /players/{riotid} endpoint
func UpdatePlayer(w http.ResponseWriter, r *http.Request) {

	// sets the "Content-Type" header of the HTTP response to JSON format.
	w.Header().Set(CONTENT_TYPE, JSON)

	// create a new pointer to an empty Player, then parse the request body into
	// that Player struct
	updatedPlayer := &models.Player{}
	utils.ParseBody(r, updatedPlayer)

	// extract riotid from the route pattern
	params := mux.Vars(r)
	riotID := params[RIOT_ID]

	// fetch the existing player in the database with the corresponding RiotID
	// and store that data in existingPlayer (type Player) and db (type gorm.DB)
	existingPlayer, db := models.GetPlayerByID(riotID)

	// update the existing player's fields with fields from the updatedPlayer.
	// if the request body contains only the fields the client wishes to update
	// and leaves all other fields empty, this code prevents the existing
	// player's data from being overwritten with empty data.
	if updatedPlayer.IRLName != "" {
		existingPlayer.IRLName = updatedPlayer.IRLName
	}
	if updatedPlayer.Team != "" {
		existingPlayer.Team = updatedPlayer.Team
	}
	if updatedPlayer.Rank != "" {
		existingPlayer.Rank = updatedPlayer.Rank
	}
	if updatedPlayer.Role != "" {
		existingPlayer.Role = updatedPlayer.Role
	}
	if updatedPlayer.Main != "" {
		existingPlayer.Main = updatedPlayer.Main
	}
	if updatedPlayer.ACS != 0 {
		existingPlayer.ACS = updatedPlayer.ACS
	}
	if updatedPlayer.KDR != 0 {
		existingPlayer.KDR = updatedPlayer.KDR
	}
	if updatedPlayer.DamagePerRound != 0 {
		existingPlayer.DamagePerRound = updatedPlayer.DamagePerRound
	}
	if updatedPlayer.HS != 0 {
		existingPlayer.HS = updatedPlayer.HS
	}

	// save the newly updated existingPlayer data to the database
	db.Save(existingPlayer)

	// encode existingPlayer into JSON format and write to w.
	utils.Write(w, existingPlayer)
}
