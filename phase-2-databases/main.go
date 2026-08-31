package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

/* Songs table
- id, name, artist, year
*/

func createSongsTable(db *sql.DB) {
	query := `create table if not exists songs (
		id serial primary key,
		name varchar(255) not null,
		artist varchar(255) not null,
		year integer not null
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	connStr := "host=192.168.2.18 port=5432 user=app1_user password=1010 dbname=app1_db sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Connection failed my Guy: ", err)
	}

	fmt.Printf("Connected successfully to PostgreSQL! 🐘🚀\n")

	createSongsTable(db)
}
