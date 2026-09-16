package database

import (
	"database/sql"
	"mymodule/models"
)

func ReadSongs(db *sql.DB) ([]models.Song, error) {
	var data []models.Song
	rows, err := db.Query("SELECT * FROM songs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song models.Song
		if err := rows.Scan(&song.ID, &song.Name, &song.Artist, &song.Year); err != nil {
			return nil, err
		}
		data = append(data, song)
	}
	return data, nil
}
