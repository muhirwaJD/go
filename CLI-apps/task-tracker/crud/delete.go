package crud

import (
	"database/sql"
	"errors"
	"fmt"
	"music/models"
)

func DeleteSong(DB *sql.DB, value models.Song) (models.Song, error) {
	var deleteSong models.Song
	query := `DELETE FROM songs WHERE id = $1 RETURNING id, name, artist, year`
	err := DB.QueryRow(query, value.ID).Scan(&deleteSong.ID, &deleteSong.Name, &deleteSong.Artist, &deleteSong.Year)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return deleteSong, fmt.Errorf("song %s doesn't exist", deleteSong.Name)
		}
		return deleteSong, fmt.Errorf("Delete song: %w", err)
	}
	return deleteSong, nil
}
