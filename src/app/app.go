package main

import (
	"encoding/json"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	httpc "httpclient"
	"httpinit"
	"mbus"
	"net/http"
	"sort"
	"strconv"
	"util"
)

const NUM_CLOSEST_STOPS = 5

func init() {
	httpinit.Init()
	goji.Get("/closest-stop", closestStop)
	goji.Get("/closest-stops", closestStops)
	goji.Get("/next-bus", nextBus)
}

func closestStop(ctx web.C, w http.ResponseWriter, r *http.Request) {
	// extract latitude and longitude from URL
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

	// create HTTP client and fetch stop list
	c := httpc.Client(r)
	stops, err := mbus.StopList(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// now sort by distance
	sort.Sort(util.ByStopDistance(stops, latitude, longitude))

	// closest stop
	stop := stops[0]

	// write this stop
	blob, err := json.Marshal(stop)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(blob)
}

func closestStops(ctx web.C, w http.ResponseWriter, r *http.Request) {
	// extract latitude and longitude from URL
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

	// create HTTP client and fetch stop list
	c := httpc.Client(r)
	stops, err := mbus.StopList(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// now sort by distance
	sort.Sort(util.ByStopDistance(stops, latitude, longitude))

	// Write final set of closest stops
	blob, err := json.Marshal(stops[:NUM_CLOSEST_STOPS])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(blob)
}

func nextBus(ctx web.C, w http.ResponseWriter, r *http.Request) {
	// get stop from URL
	stopId_string := ctx.URLParams["stop"]
	stopId, err := strconv.Atoi(stopId_string)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create HTTP client and fetch next bus
	c := httpc.Client(r)
	bus, err := mbus.NextBusAtStop(c, stopId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// now write that bus
	blob, err := json.Marshal(bus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(blob)
}
