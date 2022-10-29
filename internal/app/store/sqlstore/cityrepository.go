package sqlstore

import "github.com/MrDavudov/OpenWeatherGO/internal/app/model"

type CityRepository struct {
	store *Store
}

// Create ...
func (s *CityRepository) Create(w *model.Weather) error {
	err := s.store.Open()
	if err != nil {
		return err
	}

	defer s.store.Close()

	for i := range w.DtTemp {
		_, err := s.store.db.Exec(
			"INSERT INTO temp (data, temp) VALUES ($1, $2)",
			w.DtTemp[i].Dt,
			w.DtTemp[i].Temp)
		if err != nil {
			return err
		}
	}

	return nil
}

// Update ...
func (s *CityRepository) Update(w *model.Weather) error {
	err := s.store.Open()
	if err != nil {
		return err
	}

	defer s.store.Close()

	for i := range w.DtTemp {
		_, err := s.store.db.Exec(
			"UPDATE temp SET temp = $1 where data = $2 AND id = $3",
			w.DtTemp[i].Temp,
			w.DtTemp[i].Dt,
			w.ID)
		if err != nil {
			return err
		}
	}

	return nil
}