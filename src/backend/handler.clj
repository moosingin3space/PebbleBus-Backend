(ns backend.handler
  (:require [compojure.core :refer :all]
            [compojure.route :as route]
            [ring.middleware.defaults :refer [wrap-defaults site-defaults]]
            [backend.mbus :as mbus]
            [backend.util :as util]
            [clojure.math.numeric-tower :as math]
            [clojure.data.json :as json]))

(defroutes app-routes
  (GET "/closest-stop" {{lat :lat lon :lon} :params}
       (let [stops-list (mbus/stop-list)]
         (json/write-str 
           (mbus/mbus-to-std-json
             (util/find-closest (read-string lat) (read-string lon) stops-list)))))
  (GET "/closest-stops" {{lat :lat lon :lon} :params}
       (let [stops-list (mbus/stop-list)]
         (json/write-str
           (take 5 
                 (util/sort-closest (read-string lat) (read-string lon) (map mbus/mbus-to-std-json stops-list))))))
  (GET "/next-bus" {{stop-id :stop} :params}
       (let [etas (mbus/eta-list stop-id)
             routes (mbus/route-list)
             buses (mbus/bus-list)
             {route-id :route t :avg bus-id :bus_id} (first etas)
             route (first (filter (fn [{id :id :as all}] (= id route-id)) routes))
             {route-name :name} route]
         (json/write-str 
           {:name route-name
            :time t})))
  (route/not-found "Not Found"))

(def app
  (wrap-defaults app-routes site-defaults))
