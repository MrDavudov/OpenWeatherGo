package store

type Store interface {
	City() CityRepository
}