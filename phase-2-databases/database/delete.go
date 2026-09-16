package database

import (
	"database/sql"
	"fmt"
)

func DeleteSong(db *sql.DB, value int) error {
	var id int

	err := db.QueryRow("SELECT id from songs where id=$1", value).Scan(&id)

	if err == sql.ErrNoRows {
		return fmt.Errorf("Song not found")
	}

	row, err := db.Exec("DELETE FROM songs WHERE id = $1;", value)
	if err != nil {
		return err
	}
	count, err := row.RowsAffected()
	if count != 1 {
		totalChange := fmt.Errorf("expected 1 row deleted, got %d", count)
		return totalChange
	}

	return nil
}
