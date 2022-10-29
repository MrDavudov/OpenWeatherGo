package sqlstore

import (
	"database/sql"

	"github.com/MrDavudov/OpenWeatherGO/internal/app/store"
)

type Store struct {
	db				*sql.DB
	cityRepository	*CityRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Open ...
func (s *Store) Open() error {
	connStr := "user=admin password=admin dbname=postgres_db sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return err
    }

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

// City
func (s *Store) City() store.CityRepository {
	if s.cityRepository != nil {
		return s.cityRepository
	}

	s.cityRepository = &CityRepository{
		store: s,
	}

	return s.cityRepository
}