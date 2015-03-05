(defproject backend "0.1.0-SNAPSHOT"
  :description "FIXME: write description"
  :url "http://example.com/FIXME"
  :min-lein-version "2.0.0"
  :dependencies [[org.clojure/clojure "1.6.0"]
                 [compojure "1.3.1"]
                 [ring/ring-defaults "0.1.2"]
                 [clj-http "1.0.0"]
                 [org.clojure/data.json "0.2.5"]
                 [org.clojure/math.numeric-tower "0.0.4"]
                 [environ "0.5.0"]
                 [ring/ring-jetty-adapter "1.2.2"]]
  :plugins [[lein-ring "0.8.13"]
            [environ/environ.lein "0.2.1"]]
  :ring {:handler backend.handler/app}
  :hooks [environ.leiningen.hooks]
  :uberjar-name "pebblebus.jar"
  :profiles 
  {:dev {:dependencies [[javax.servlet/servlet-api "2.5"]
                        [ring-mock "0.1.5"]]}
   :production {:env {:production true}}})
