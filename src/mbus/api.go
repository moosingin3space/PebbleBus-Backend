package mbus

import (
	"data"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Route struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

const API_ENDPOINT = "http://mbus.doublemap.com/map/v2"

func StopList(cl *http.Client) (stops []data.Stop, err error) {
	stopsUrl := API_ENDPOINT + "/stops"
	resp, err := cl.Get(stopsUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &stops)
	return
}

func NextBusAtStop(cl *http.Client, stopId int) (bus data.Bus, err error) {
	etasUrl := API_ENDPOINT + "/eta?stop=" + strconv.Itoa(stopId)
	routesUrl := API_ENDPOINT + "/routes"
	etaResp, err := cl.Get(etasUrl)
	if err != nil {
		return
	}
	routeResp, err := cl.Get(routesUrl)
	if err != nil {
		return
	}
	defer etaResp.Body.Close()
	defer routeResp.Body.Close()
	etaBody, err := ioutil.ReadAll(etaResp.Body)
	if err != nil {
		return
	}
	routeBody, err := ioutil.ReadAll(routeResp.Body)
	if err != nil {
		return
	}
	// read etas
	var etas map[string]interface{}
	err = json.Unmarshal(etaBody, &etas)
	if err != nil {
		return
	}
	innerEtas := etas["etas"].(map[string]interface{})[strconv.Itoa(stopId)].(map[string]interface{})["etas"].([]interface{})
	if len(innerEtas) > 0 {
		first_eta := innerEtas[0].(map[string]interface{})
		routeId := int(first_eta["route"].(float64))

		// read routes
		var routes []Route
		err = json.Unmarshal(routeBody, &routes)
		if err != nil {
			return
		}
		for _, r := range routes {
			if r.Id == routeId {
				time, ok := first_eta["avg"].(float64)
				if ok {
					// our route
					bus = data.Bus{
						Text: r.Name,
						Time: int(time),
					}
					return
				} else {
					err = errors.New("invalid type")
					return
				}
			}
		}
	}
	err = errors.New("no bus found")
	return
}
