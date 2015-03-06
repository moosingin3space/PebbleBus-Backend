(ns backend.mbus
  (:require [clojure.data.json :as json]
            [clj-http.client :as client]))

(def api-endpoint "http://mbus.doublemap.com/map/v2")

(defn stop-list [] 
  (let [stops-url (str api-endpoint "/stops")
        {stops-raw :body} (client/get stops-url)
        stops-response (json/read-str stops-raw :key-fn keyword)]
    stops-response))

(defn eta-list [stop-id]
  (let [eta-url (str api-endpoint "/eta?stop=" stop-id)
        {etas-raw :body} (client/get eta-url)
        {{{etas-response :etas} (keyword (str stop-id))} :etas} (json/read-str etas-raw :key-fn keyword)]
    etas-response))

(defn route-list []
  (let [routes-url (str api-endpoint "/routes")
        {routes-raw :body} (client/get routes-url)
        routes-resp (json/read-str routes-raw :key-fn keyword)]
    routes-resp))

(defn bus-list []
  (let [buses-url (str api-endpoint "/buses")
        {buses-raw :body} (client/get buses-url)
        buses-response (json/read-str buses-raw :key-fn keyword)]
    buses-response))

(defn mbus-to-std-json [{id :id nm :name lat :lat lon :lon}]
  {:id id :name nm :lat lat :lon lon})
