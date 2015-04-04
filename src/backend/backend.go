package backend

import (
	"data"
	"net/http"
)

type Creator func(cl *http.Client) (Backend, error)

type Backend interface {
	ClosestStops(lat, lon float64, num int) ([]data.Stop, error)
}
