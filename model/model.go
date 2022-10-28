package model

import (
	"encoding/json"
	"os"
)

type Weather struct {
	Name     string
	Lat      float64
	Lon      float64
	Country  string
	DataTemp []DtTemp
}

type DtTemp struct {
	dt   string
	temp float64
}

func JsonSave(m *Weather) error {
	const jsonFile = "./db.json"

	type city struct {
		id		int
		name	string
		lat		float64
		lon		float64
		country	string
	}

	// Проверка существует ли такой файл
	_, err := os.Stat(jsonFile)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create("db.json")
			if err != nil {
				return err
			}

			err = os.WriteFile(jsonFile, []byte("[]"), 0666)
			if err != nil {
				return err
			}
		}
	}

	// Чтения файла для преобразование
	rawDataIn, err := os.ReadFile(jsonFile)
	if err != nil {
		return err
	}

	setting := []city{}

	err = json.Unmarshal(rawDataIn, &setting)
	if err != nil {
		return err
	}
	
	// добавления города если его нет
	for i := range setting {
		if setting[i].name == m.Name {
			return nil
		}
	}

	setting = append(setting, city{
		id: len(setting)+1,
		name: m.Name,
		lat: m.Lat,
		lon: m.Lon,
		country: m.Country,
	})

	return nil
}