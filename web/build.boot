(set-env!
 :source-paths #{"src/cljs"}
 :resource-paths #{"html"}

 :dependencies '[
                 [org.clojure/clojure "1.8.0"]
                 [org.clojure/clojurescript "1.8.40"]
                 [adzerk/boot-cljs "1.7.228-1"]
                 [pandeiro/boot-http "0.7.3"]
                 [adzerk/boot-reload "0.4.7"]
                 [re-frame "0.7.0"]
                 [reagent "0.6.0-alpha"]
                 ])

(require '[adzerk.boot-cljs :refer [cljs]]
         '[pandeiro.boot-http :refer [serve]]
         '[adzerk.boot-reload :refer [reload]])

(deftask dev
  []
  (comp
   (serve :dir "target")
   (watch)
   (reload :on-jsload 'us.nicot.core/run)
   (cljs)
   (target :dir #{"target"})))
