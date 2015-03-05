(ns backend.main
  (:require [ring.adapter.jetty :as jetty]
            [backend.handler :as handler]
            [compojure.handler :refer [site]]
            [environ.core :refer [env]]))

(defn -main [& [port]]
  (let [port (Integer. (or port (env :port) 5000))]
    (jetty/run-jetty (site #'handler/app) {:port port :join? false})))

