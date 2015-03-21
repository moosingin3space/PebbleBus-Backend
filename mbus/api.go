package mbus

import (
	"appengine"
	"appengine/urlfetch"
	"data"
	"encoding/json"
	"io/ioutil"
)

const API_ENDPOINT = "http://mbus.doublemap.com/map/v2"

func StopList(ctx appengine.Context) (stops []data.Stop, err error) {
	stopsUrl := API_ENDPOINT + "/stops"
	client := urlfetch.Client(ctx)
	resp, err := client.Get(stopsUrl)
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
