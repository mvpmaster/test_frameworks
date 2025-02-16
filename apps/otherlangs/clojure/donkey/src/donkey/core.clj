;; (ns com.appsflyer.sample-app
;;   (:require [com.appsflyer.donkey.core :refer [create-donkey create-server]]
;;           [com.appsflyer.donkey.server :refer [start]]
;;           [com.appsflyer.donkey.result :refer [on-success]])



;;   (def ^Donkey donkey-core (create-donkey))

(ns donkey.core
  (:require [com.appsflyer.donkey.core :refer [create-donkey create-server]]
          [com.appsflyer.donkey.server :refer [start]]
          [com.appsflyer.donkey.result :refer [on-success]]
          [clojure.data.json :as json]
          [cheshire.core :refer :all]
          ;;[ring.middleware.json :refer [wrap-json-response]]
        ))

;; https://github.com/ring-clojure/ring-json
;; https://stackoverflow.com/questions/3436216/how-to-map-clojure-code-to-and-from-json
;; https://clojuredocs.org/clojure.core/load-file
;; https://stackoverflow.com/questions/15660066/how-to-read-json-file-into-clojure-defrecord-to-be-searched-later
;; https://ericnormand.me/mini-guide/json-serialization-api-clojure


(def myjson (json/read-str (slurp "test.json")))

(println myjson)

;;(def ^Donkey donkey-core (create-donkey))

;;  (def donkey-core (create-donkey {:event-loops 4}))

;; (ns test.core
;;   (:gen-class))

(def test (apply str (repeat 5000 "test")))  ;;"test")
(def simple (apply str (repeat 2 "test")))  ;;"test")

(->
  (create-donkey)
  (create-server {:port   8080
                  :routes [
                    {
                    :path         "/"
                    :handler (fn [_request respond _raise]
                                       (respond {:body simple ;; (apply str (repeat 3 "str")) ;;"Hello, world!"
                                                }))},
                    {
                    :path         "/test"
                    :handler (fn [_request respond _raise]
                                       (respond {:body test ;; (apply str (repeat 3 "str")) ;;"Hello, world!"
                                                }))}
                    {
                    :path         "/json"
                    :handler (fn [_request respond _raise]
                                       (respond {:body (generate-string myjson {:escape-non-ascii false}) ;;(json/write-str myjson) ;; (apply str (repeat 3 "str")) ;;"Hello, world!"
                                                }))}
                                          
                                                
                      ]})
  start
  (on-success (fn [_] (println "Server started listening on port 8080"))))

;; (defn -main
;;   "I don't do a whole lot ... yet."
;;   [& args]
;;   (println "Hello, World!"))


;; (defn -main
;;   "I don't do a whole lot ... yet."
;;   [& args]
;;   (println "Hello, World!"))
