package crud

import (
	"database/sql"
	"fmt"
	"music/models"
)

func ReadSongs(DB *sql.DB) ([]models.Song, error) {
	query := `SELECT id, name, artist, year FROM songs ORDER BY id`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("list songs: %w", err)
	}

	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var s models.Song
		err := rows.Scan(&s.ID, &s.Name, &s.Artist, &s.Year)
		if err != nil {
			return nil, fmt.Errorf("scan song: %w", err)
		}
		songs = append(songs, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate songs: %w", err)
	}

	return songs, nil
}
