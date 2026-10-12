package crud

import (
	"database/sql"
	"errors"
	"fmt"
	"music/models"

	"github.com/lib/pq"
	"github.com/lib/pq/pqerror"
)

var ErrSongNotFound = errors.New("song not found")

func Update(DB *sql.DB, value models.Song) (models.Song, error) {
	var updated models.Song
	query := `UPDATE songs SET name=$1, artist=$2, year=$3 WHERE id=$4 RETURNING name, artist, year, id`

	err := DB.QueryRow(query, value.Name, value.Artist, value.Year, value.ID).Scan(&updated.Name, &updated.Artist, &updated.Year, &updated.ID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Song{}, fmt.Errorf("id %d: %w", value.ID, ErrSongNotFound)
		}
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == pqerror.UniqueViolation {
			return models.Song{}, fmt.Errorf("%q by %s: %w", value.Name, value.Artist, ErrDuplicateSong)
		}
		return models.Song{}, fmt.Errorf("update song %d: %w", value.ID, err)
	}

	return updated, nil
}
