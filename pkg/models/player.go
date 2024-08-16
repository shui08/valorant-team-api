// this code contains the player struct, as well as logic for directly
// interacting with the database.
package models

import (
	"github.com/shui08/valorant-team-api/pkg/config"
	"gorm.io/gorm"
)

const (
	RIOT_COND = "riot_id = ?"
)

// used to refer to the database connection
var db *gorm.DB

// the model for what data a player instance will hold. the JSON tags specify
// how the data will be displayed when marshaled to JSON.
type Player struct {
	RiotID         string  `json:"riotid"` // must be entered in form: USER-TAG. ex: John-123
	IRLName        string  `json:"irlname"`
	Team           string  `json:"team"`
	Rank           string  `json:"rank"`
	Role           string  `json:"role"`
	Main           string  `json:"main"`
	ACS            float64 `json:"acs"`
	KDR            float64 `json:"kdr"`
	DamagePerRound float64 `json:"dpr"`
	HS             float64 `json:"hs"`
}

// init() functions are run automatically when the package loads.
func init() {

	// create a database connection
	config.Connect()

	// store the connection instance in db
	db = config.GetDB()

	// auto migrates the Player struct to a database table. this means that
	// the Player struct will automatically be represented in a database table
	// if it does not already exist, and if it does, the table will update to
	// include any new fields that may have been created.
	err := db.AutoMigrate(&Player{})

	// if the auto migration fails, panic will be called.
	if err != nil {
		panic(err)
	}
}

// this function "creates" a player by adding the player to the database. it
// has a receiver that points to a Player instance, and returns that pointer
func (player *Player) AddPlayer() *Player {

	// this creates a new record for the player data inside the database.
	db.Create(player)
	return player
}

// this function will query for all players in the database and return those
// players in a slice.
func GetAllPlayers() []Player {
	// creating a slice to store the players in
	var players []Player

	// db.Find() queries the database for all records associated with the struct
	// that is passed in. in this case, it fetches all records corresponding to
	// Player. it then populates the struct with those records (in other words,
	// it populates the Players slice with all the players found in the
	// database)
	db.Find(&players)

	// return the populated Players slice
	return players
}

// this function retrieves a player by their Riot ID.
func GetPlayerByID(RiotID string) (*Player, *gorm.DB) {

	// declaring a Player instance whose fields we will populate with the query
	// results from db.Find
	var player Player

	// querying the database for a player whose riot id matches the specified
	// RiotID argument. the player instance's fields will then be populated with
	// the resulting data. we also initialize a new local db variable, which
	// is of type gorm.DB and holds the results of the query.
	db := db.Find(&player, RIOT_COND, RiotID)

	// return the newly populated player instance and the local db variable.
	// the db variable will be useful for controllers.UpdatePlayer, when we have
	// to update a record in the database.
	return &player, db
}

// this function deletes the player specified by the RiotID argument from the
// database. it then returns that player instance.
func DeletePlayer(RiotID string) Player {

	// declaring a Player instance whose fields we will populate with the query
	// results from db.Find
	var player Player

	// querying the database for a player whose riot id matches the specified
	// RiotID argument. the player instance's fields will then be populated with
	// the resulting data.
	db.Find(&player, RIOT_COND, RiotID)

	// once again querying for a player whose riot id matches the RiotID
	// argument, and then deleting that player from the database. (i query twice
	// because i honestly am not sure whether the query from the previous line
	// of code carries over, and i do not want to accidentally perform a batch
	// deletion).
	db.Where(RIOT_COND, RiotID).Delete(&player)

	// return the deleted player
	return player
}

// this function deletes all players from the database and should be used with
// caution.
func DeleteAll() ([]Player, error) {
	// creating a slice to store the players in
	var players []Player

	// querying the database for all players and storing them in `players`
	db.Find(&players)

	// enabling global updates, which allows large batch deletions, and then
	// deleting all player records from the database
	deleteResult := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Player{})

	// if the deletion fails, return an error
	if deleteResult.Error != nil {
		return nil, deleteResult.Error
	}

	// otherwise, return the slice of players
	return players, nil
}
