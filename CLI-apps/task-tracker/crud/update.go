package crud

import (
	"database/sql"
	"fmt"
	"music/models"
)

func Update(DB *sql.DB, value models.Song) error {
	fmt.Print("We're updating our freaking song")
	name := value.Name
	year := value.Year
	artist := value.Artist
	id := value.ID

	row, err := DB.Exec("UPDATE songs SET name=$1, artist=$2, year=$3 WHERE id=$4", name, artist, year, id)
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
