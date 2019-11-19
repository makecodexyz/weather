package models

type CityWeather struct {
	City        string `json:"city"`
	UpdateTime  string `json:"update_time"`
	Weather     string `json:"weather"`
	Temperature string `json:"temperature"`
	Wind        string `json:"wind"`
}
