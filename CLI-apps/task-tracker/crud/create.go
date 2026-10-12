package crud

import (
	"database/sql"
	"errors"
	"fmt"
	"music/models"

	"github.com/lib/pq"
	"github.com/lib/pq/pqerror"
)

var ErrDuplicateSong = errors.New("song already exists")

func Create(DB *sql.DB, value models.Song) (models.Song, error) {
	var saved models.Song

	query := `INSERT INTO songs (name, artist, year)
			VALUES ($1, $2, $3)
			RETURNING id, name, artist, year;`

	err := DB.QueryRow(query, value.Name, value.Artist, value.Year).Scan(&saved.ID, &saved.Name, &saved.Artist, &saved.Year)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == pqerror.UniqueViolation {
			return models.Song{}, fmt.Errorf("%q by %s: %w", value.Name, value.Artist, ErrDuplicateSong)
		}
		return models.Song{}, fmt.Errorf("create song %q: %w", value.Name, err)
	}
	return saved, nil
}
