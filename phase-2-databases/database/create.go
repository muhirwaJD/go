package database

import (
	"database/sql"
	"fmt"
	"mymodule/models"
)

func InsertSong(db *sql.DB, value models.Song) error {
	name := value.Name
	artist := value.Artist
	year := value.Year

	exec, err := db.Exec("INSERT INTO songs (name, artist, year)"+
		"VALUES ($1, $2, $3);", name, artist, year)
	if err != nil {
		return err
	}

	results, err := exec.RowsAffected()
	if err != nil {
		return err
	}

	if results != 1 {
		return fmt.Errorf("expected 1 row affected, got %d", results)
	}

	return nil
}
