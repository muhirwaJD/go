package crud

import (
	"database/sql"
	"fmt"
	"music/models"
)

func Create(DB *sql.DB, value models.Song) error {
	fmt.Println("Hello We're going to create a song in the DB")

	name := value.Name
	year := value.Year
	artist := value.Artist

	exec, err := DB.Exec("INSERT INTO songs (name, artist, year)"+
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
