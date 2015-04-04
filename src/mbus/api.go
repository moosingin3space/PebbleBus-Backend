package mbus

import (
	"backend"
	"data"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"util"
)

type Route struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

const API_ENDPOINT = "http://mbus.doublemap.com/map/v2"

type Service struct {
	stops  []data.Stop
	routes []Route
	client *http.Client
}

func InitService(cl *http.Client) (svc backend.Backend, err error) {
	var stops []data.Stop
	var routes []Route
	stopsUrl := API_ENDPOINT + "/stops"
	routesUrl := API_ENDPOINT + "/routes"

	// Get Stops array
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
	if err != nil {
		return
	}

	// Get Routes array
	resp, err = cl.Get(routesUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &routes)
	if err != nil {
		return
	}
	svc = Service{
		stops:  stops,
		routes: routes,
		client: cl,
	}
	return
}

func (svc Service) ClosestStops(lat, lon float64, num int) (stops []data.Stop, err error) {
	myStops := svc.stops[:]
	sort.Sort(util.ByStopDistance(myStops, lat, lon))
	stops = make([]data.Stop, num)
	j := 0
	for _, stop := range myStops {
		if j >= num {
			break
		}
		etasUrl := API_ENDPOINT + "/eta?stop=" + strconv.Itoa(stop.Id)
		etaResp, err := svc.client.Get(etasUrl)
		if err != nil {
			return nil, err
		}
		defer etaResp.Body.Close()
		etaBody, err := ioutil.ReadAll(etaResp.Body)
		if err != nil {
			return nil, err
		}
		// read etas
		var etas map[string]interface{}
		err = json.Unmarshal(etaBody, &etas)
		if err != nil {
			return nil, err
		}
		innerEtas := etas["etas"].(map[string]interface{})[strconv.Itoa(stop.Id)].(map[string]interface{})["etas"].([]interface{})
		if len(innerEtas) > 0 {
			first_eta := innerEtas[0].(map[string]interface{})
			routeId := int(first_eta["route"].(float64))
			for _, r := range svc.routes {
				if r.Id == routeId {
					time, ok := first_eta["avg"].(float64)
					if ok {
						// our route
						bus := data.Bus{
							Text: r.Name,
							Time: int(time),
						}
						stop.Buses = []data.Bus{bus}
						stops[j] = stop
						j++
					} else {
						err = errors.New("invalid type")
						return nil, err
					}
				}
			}
		}
	}
	return stops, nil
}
