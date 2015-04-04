package main

import (
	"backend"
	"data"
	"encoding/json"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	httpc "httpclient"
	"httpinit"
	"mbus"
	"net/http"
	"regexp"
	"strconv"
)

const NUM_CLOSEST_STOPS = 5

var services []backend.Creator = []backend.Creator{
	mbus.InitService,
}

func init() {
	httpinit.Init()
	re := regexp.MustCompile(`^/closest-stops/@(?P<lat>[\d.-]+),(?P<lon>[\d.-]+)$`)
	goji.Get(re, closest)
}

func closest(ctx web.C, w http.ResponseWriter, r *http.Request) {
	lat_string := ctx.URLParams["lat"]
	lon_string := ctx.URLParams["lon"]
	latitude, err := strconv.ParseFloat(lat_string, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	longitude, err := strconv.ParseFloat(lon_string, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c := httpc.Client(r)
	var stops []data.Stop
	// Iterate over all the services
	for _, fact := range services {
		service, err := fact(c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		stops, err = service.ClosestStops(latitude, longitude, NUM_CLOSEST_STOPS)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	// now write that list of stops
	blob, err := json.Marshal(stops)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(blob)
}
