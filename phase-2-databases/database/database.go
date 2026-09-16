package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.2.18"
	port     = 5432
	user     = "app1_user"
	password = "1010"
	dbname   = "app1_db"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SetupConstraints(db *sql.DB) error {
	// Check if constraint already exists
	var exists bool
	err := db.QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM information_schema.table_constraints
			WHERE constraint_name = 'unique_song'
			AND table_name = 'songs'
		)
	`).Scan(&exists)
	if err != nil {
		return err
	}

	// If constraint doesn't exist, create it
	if !exists {
		_, err := db.Exec(`
			ALTER TABLE songs
			ADD CONSTRAINT unique_song UNIQUE(name, artist)
		`)
		if err != nil {
			return err
		}
	}

	return nil
}
