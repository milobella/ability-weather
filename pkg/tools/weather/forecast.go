package weather

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

func (wd *WeatherData) UnmarshalJSON(b []byte) (err error) {
	var rawStrings map[string]string

	if err = json.Unmarshal(b, &rawStrings); err != nil {
		return err
	}

	for k, v := range rawStrings {
		if strings.ToLower(k) == "temperature" {
			if wd.Temperature, err = strconv.ParseFloat(v, 32); err != nil {
				return err
			}
		}

		if strings.ToLower(k) == "timestamp" {
			if wd.Timestamp, err = time.Parse(time.RFC3339, v); err != nil {
				return err
			}
		}

		if strings.ToLower(k) == "weather" {
			wd.Weather = v
		}
	}
	return
}

type Forecast struct {
	Forecast []WeatherData
	Count    int
}

type WeatherData struct {
	Temperature float64   `json:"temperature,string"`
	Timestamp   time.Time `json:"timestamp,string"`
	Weather     string    `json:"weather,string"`
}
