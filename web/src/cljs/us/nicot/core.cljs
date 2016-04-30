(ns us.nicot.core
  (:require [reagent.core :as reagent]))

(defn link [text href]
  [:a {:href href
       :style {:margin "0px 10px"
               :text-decoration "none"}}
   text])

(defn links []
  [:span
   {:style
    {:margin-left "auto"}}
   (link "Github" "https://github.com/nicot")
   (link "LinkedIn" "https://example.com/TODO")
   (link "Contact" "/contact")
   ])

(defn header []
  [:div
   {:style {:font-family "Roboto"
            :font-size "24px"
            :border-bottom-style "solid"
            :border-width "1px"
            :padding "5px"
            :display "flex"}}
   [link "Nico Tonozzi" "/"]
   [links]])

(defn page []
  (fn []
    [:div
     {:style {:width "80vw"
              :max-width "1000px"
              :min-width "300px"
              :margin "auto"
              :margin-top "20px"}}
     [header]]))

(defn ^:export run []
  (reagent/render [page] (.getElementById js/document "app")))
