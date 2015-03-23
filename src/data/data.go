package data

type Stop struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type Bus struct {
	Text string `json:"name"`
	Time int    `json:"time"`
}
