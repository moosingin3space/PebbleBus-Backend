package mbus

import (
	"data"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
