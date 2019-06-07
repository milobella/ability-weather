package weather

import (
	"encoding/json"
	"strings"
	"time"
)

func (wd *WeatherData) UnmarshalJSON(b []byte) (err error) {
	var rawStrings map[string]interface{}

	if err = json.Unmarshal(b, &rawStrings); err != nil {
		return err
	}

	for k, v := range rawStrings {
		if strings.ToLower(k) == "temperature" {
			wd.Temperature = v.(float64)
		}

		if strings.ToLower(k) == "timestamp" {
			wd.Timestamp = time.Unix(int64(v.(float64)), 0)
		}

		if strings.ToLower(k) == "weather" {
			wd.Weather = v.(string)
		}
	}
	return
}

type Forecast struct {
	Forecast []WeatherData
	Count    int
}

type WeatherData struct {
	Temperature float64   `json:"temperature,number"`
	Timestamp   time.Time `json:"timestamp,string"`
	Weather     string    `json:"weather,string"`
}
