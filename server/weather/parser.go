package weather

type WeatherResponse struct {
	Main struct {
		FeelsLike float64 `json:"feels_like"`
		Temp      float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}
