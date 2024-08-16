// this code aims to create a connection to a MySQL database as well as
// a db (database) variable that can be used throughout the program to refer
// to the database connection (use GetDB to access it).
package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// declaring db, which will be used to refer to the database connection
var db *gorm.DB

// init() functions are run automatically when the package loads, this is simply
// to load any environment variables
func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

// this function is what creates a connection to the database, and we will use
// gorm to do so. the reason for using gorm is so that we can interact with the
// database using Go, rather than SQL. mysql.Open() takes in a database
// connection string and returns a MySQL driver instance, which gorm can use
// to connect with the database. gorm.Open() takes in that driver instance as
// well as a config struct, and actually creates the database connection, which
// is stored in `data` as a *gorm.DB instance. it also returns an error, and if
// that error is non-nil, we will call panic. we then set db, the package-wide
// connection instance variable, equal to data.
func Connect() {
	dsn := os.Getenv("DB_DSN")
	data, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = data
}

// this function is simply for accessing the db variable outside of the package.
func GetDB() *gorm.DB {
	return db
}
