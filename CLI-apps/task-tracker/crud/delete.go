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
			return models.Song{}, fmt.Errorf("id %d: %w", value.ID, ErrSongNotFound)
		}
		return models.Song{}, fmt.Errorf("delete song %d: %w", value.ID, err)
	}
	return deleteSong, nil
}
