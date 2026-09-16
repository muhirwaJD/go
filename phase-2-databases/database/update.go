package database

import (
	"database/sql"
	"fmt"
	"mymodule/models"
)

func UpdateSong(db *sql.DB, value models.Song) error {
	id := value.ID
	name := value.Name
	artist := value.Artist
	year := value.Year

	row, err := db.Exec("UPDATE songs SET name=$1, artist=$2, year=$3 WHERE id=$4", name, artist, year, id)
	if err != nil {
		return err
	}

	count, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		totalChange := fmt.Errorf("expected 1 row updated, got %d", count)
		return totalChange
	}

	return nil
}
