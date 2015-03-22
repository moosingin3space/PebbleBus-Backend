package util

import (
	"data"
	"math"
)

func sqrDist(lat1, lon1, lat2, lon2 float64) float64 {
	return math.Pow(lat2-lat1, 2) + math.Pow(lon2-lon1, 2)
}

type sqrDistFunc func(float64, float64) float64

type byStopDistance struct {
	stops     []data.Stop
	latitude  float64
	longitude float64
	myDist    sqrDistFunc
}

func ByStopDistance(stops []data.Stop, lat float64, lon float64) byStopDistance {
	return byStopDistance{
		stops:     stops,
		latitude:  lat,
		longitude: lon,
		myDist: func(lat1, lon1 float64) float64 {
			return sqrDist(lat1, lon1, lat, lon)
		},
	}
}

func (a byStopDistance) Len() int {
	return len(a.stops)
}

func (a byStopDistance) Swap(i, j int) {
	a.stops[i], a.stops[j] = a.stops[j], a.stops[i]
}

func (a byStopDistance) Less(i, j int) bool {
	return a.myDist(a.stops[i].Latitude, a.stops[i].Longitude) < a.myDist(a.stops[j].Latitude, a.stops[j].Longitude)
}
