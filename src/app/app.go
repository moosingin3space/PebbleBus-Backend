package main

import (
	"encoding/json"
	"fmt"
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
	w.WriteHeader(500)
	fmt.Fprint(w, "Not implemented")
}

func closestStops(w http.ResponseWriter, r *http.Request) {
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
	c := httpc.Client(r)
	stops, err := mbus.StopList(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// now sort by distance
	sort.Sort(util.ByStopDistance(stops, latitude, longitude))

	blob, err := json.Marshal(stops[:NUM_CLOSEST_STOPS])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(blob)
}

func nextBus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	fmt.Fprint(w, "Not implemented")
}
