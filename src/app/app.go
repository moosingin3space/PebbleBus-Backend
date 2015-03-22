package main

import (
	"encoding/json"
	"fmt"
	httpc "httpclient"
	"mbus"
	"net/http"
)

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
	c := httpc.Client(r)
	stops, err := mbus.StopList(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	blob, err := json.Marshal(stops)
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
