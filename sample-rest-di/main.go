package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/sanjushahgupta/lessons/sample-rest-di/application"
	"github.com/sanjushahgupta/lessons/sample-rest-di/handler"
	"github.com/sanjushahgupta/lessons/sample-rest-di/store/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

// TODO db name and port are hard coded, should read from env
// DETAILS env set up
func main() {
	//  TODO db name and port are hard coded, should read from env
	dbname := os.Getenv("dbname")
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	initDB(db)

	// instantiate necessary layers via dependency injection
	store := sqlite.NewStore(db)
	logic := application.NewLogic(store)
	rest := handler.NewREST(logic)
	Dbport := os.Getenv("port")

	// set up routes
	router := httprouter.New()
	router = rest.SetupUserRoutes(router)
	log.Fatal(http.ListenAndServe(Dbport, router))
	// http.ListenAndServe(":8080", router)
}

// initDB sets up initial database structure in user.db
func initDB(db *sql.DB) {
	// TODO SQL create statements should exists in a .sql file for ease of migration
	// DETAILS sql for create statements
	// need to change code
	// change code == bad -> might break
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS user (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, first_name TEXT, last_name TEXT, email TEXT);
	`

	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}

	file, err := ioutil.ReadFile("todo.sql")
	if err != nil {
		log.Fatalf("%q: %s\n", err, file)
	}
	todofile := strings.Split(string(file), ";")
	for _, todo := range todofile {
		_, err := db.Exec(todo)
		if err != nil {
			log.Fatalf("%q: %s\n", err, todo)
		}

	}

}
