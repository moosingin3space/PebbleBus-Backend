package main

import (
	"encoding/json"
	httpc "httpclient"
	"mbus"
	"net/http"
	"sort"
	"strconv"
	"util"
)

const NUM_CLOSEST_STOPS = 5

func init() {
	http.HandleFunc("/closest-stop", closestStop)
	http.HandleFunc("/closest-stops", closestStops)
	http.HandleFunc("/next-bus", nextBus)
}

func closestStop(w http.ResponseWriter, r *http.Request) {
	// extract latitude and longitude from URL
	lat_string := r.URL.Query().Get("lat")
	lon_string := r.URL.Query().Get("lon")
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

func closestStops(w http.ResponseWriter, r *http.Request) {
	// extract latitude and longitude from URL
	lat_string := r.URL.Query().Get("lat")
	lon_string := r.URL.Query().Get("lon")
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

func nextBus(w http.ResponseWriter, r *http.Request) {
	// get stop from URL
	stopId_string := r.URL.Query().Get("stop")
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
