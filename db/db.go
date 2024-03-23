package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB // Database connection handle (initialized in InitDB)

// InitDB establishes a connection to the SQLite database "api.db".
func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		return fmt.Errorf("could not connect to DB: %w", err) // Wrap and provide context
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	if err := createTable(); err != nil {
		return fmt.Errorf("could not create events table: %w", err) // Wrap and provide context
	}

	// Defer closing the database connection on function exit
	defer func() {
		if err := DB.Close(); err != nil {
			log.Printf("Error closing database connection: %v\n", err)
		}
	}()

	return nil
}

// createTable creates the "events" table if it doesn't exist.
func createTable() error {
	creatEventsTable := `
CREATE TABLE IF NOT EXISTS events(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    dateTime DATETIME NOT NULL,
    user_id INTEGER
)
`
	_, err := DB.Exec(creatEventsTable)
	if err != nil {
		return fmt.Errorf("could not create events table: %w", err)
	}
	return nil
}
