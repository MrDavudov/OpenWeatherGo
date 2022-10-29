package store

import "github.com/MrDavudov/OpenWeatherGO/internal/app/model"

type CityRepository interface {
	Create(*model.Weather) error
	Update(*model.Weather) error
}