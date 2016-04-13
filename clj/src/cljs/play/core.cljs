(ns play.core
  (:require
    [garden.core :refer [css]]
    [reagent.core :as reagent]
    [matchbox.core :as matchbox]
    [re-frame.core :as reframe]
    ))

(enable-console-print!)

(def app-state (reagent/atom {}))

(defn canon [city]
  (clojure.string/lower-case (clojure.string/replace city #" " "")))

(defrecord City [name coords])

(def cities [(City. "San Francisco" [37.78, -122.41])
             (City. "Charlotte" [35.22, -80.84])
             (City. "Denver" [39.83, -104.99])
             (City. "New York" [40.71, -74.00])
             (City. "San Diego" [32.71, -117.16])])

(doseq [city cities]
  (matchbox/deref (matchbox/connect (str
                                     "https://publicdata-weather.firebaseio.com/"
                                     (canon (:name city)) "/currently"))
                  #(swap! app-state assoc city ^{:key city} %&)))



(defn select [city]
  (swap! app-state assoc :selected city))

(defn display [city weather] ^{:key city}
  [:tr {:on-click #(select city)}
   [:td  city]
   [:td (:summary (first weather))]
   [:td (:temperature (first weather))]
   [:td (:windSpeed (first weather)) ]])

(defn home-render [weather]
  (let [rows (doall (for [city cities]
                      (display (:name city) (get weather city))))]
    [:div.container
     [:div#map {:style {:height "600px" :width "100%"}}]
     [:h1 "What's the weather like?"]
     [:div#cities
      [:table.highlight.bordered
       [:thread [:tr
                 [:th "City"]
                 [:th "Weather"]
                 [:th "Temperature (°F)"]
                 [:th "Wind Speed (MPH)"]]]
       [:tbody rows]]]]))

(defn color [city]
  (prn (:selected @app-state))
  (if (= city (:selected @app-state)) "red" "blue"))

(defn home-did-mount [app-state]
  (let [zoom 4
        map (.setView (.map js/L "map") #js [40.80 -100.70] zoom)
        root "https://api.mapbox.com/v4/mapbox.streets/{z}/{x}/{y}.png?access_token="
        token "pk.eyJ1Ijoibmljb3QiLCJhIjoiY2lqdnkzbGQyMGRqY3VjbTVwbDNyOGcxaiJ9.-tW3kDrfp15nLw82zErsjg"]
    (.addTo
     (.tileLayer
      js/L (str root token)
      (clj->js {:maxZoom 18}))
     map)
    (doseq [city cities]
      (let [radius 100000
            marker (.circle js/L (clj->js (:coords city)) radius (clj->js {:color (color city)}))]
        (.addTo marker map)
        (.on marker "click" #(select (:name city)))))))

(defn style []
  (css
   [:body {:font-size "20px"}]
   [:h1 {:text-align "center"}]))

(defn home []
  (reagent/create-class {:reagent-render #(home-render @app-state)
                         :component-did-mount #(home-did-mount @app-state)}))

(defn ^:export main []
  (reagent/render [home]
                  (.getElementById js/document "app"))
  (set! (.-innerText (.getElementById js/document "style")) (style)))
