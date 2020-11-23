package model

type Weather struct {
	ConsolidatedWeather []ConsolidatedWeather `json:"consolidated_weather"`
	Time                string                `json:"time"`
	SunRise             string                `json:"sun_rise"`
	SunSet              string                `json:"sun_set"`
	TimezoneName        string                `json:"timezone_name"`
	Parent              Parent                `json:"parent"`
	Sources             []Source              `json:"sources"`
	Title               string                `json:"title"`
	LocationType        string                `json:"location_type"`
	Woeid               int64                 `json:"woeid"`
	LattLong            string                `json:"latt_long"`
	Timezone            string                `json:"timezone"`
}

type ConsolidatedWeather struct {
	ID                   int64   `json:"id"`
	WeatherStateName     string  `json:"weather_state_name"`
	WeatherStateAbbr     string  `json:"weather_state_abbr"`
	WindDirectionCompass string  `json:"wind_direction_compass"`
	Created              string  `json:"created"`
	ApplicableDate       string  `json:"applicable_date"`
	MinTemp              float64 `json:"min_temp"`
	MaxTemp              float64 `json:"max_temp"`
	TheTemp              float64 `json:"the_temp"`
	WindSpeed            float64 `json:"wind_speed"`
	WindDirection        float64 `json:"wind_direction"`
	AirPressure          float64 `json:"air_pressure"`
	Humidity             int64   `json:"humidity"`
	Visibility           float64 `json:"visibility"`
	Predictability       int64   `json:"predictability"`
}

type Parent struct {
	Title        string `json:"title"`
	LocationType string `json:"location_type"`
	Woeid        int64  `json:"woeid"`
	LattLong     string `json:"latt_long"`
}

type Source struct {
	Title     string `json:"title"`
	Slug      string `json:"slug"`
	URL       string `json:"url"`
	CrawlRate int64  `json:"crawl_rate"`
}
