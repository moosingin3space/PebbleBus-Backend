// +build appengine
package data

type Stop struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lon"`
}

type Bus struct {
	Text string `json:"name"`
	Time int    `json:"time"`
}
