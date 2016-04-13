(set-env!
 :resource-paths #{"src/clj"}
 :dependencies '[[org.clojure/clojure "1.8.0"]
                 [org.clojure/core.async "0.2.374"]
                 [net.async/async "0.1.0"]])

(task-options!
 jar {:main 'play.core
      :file "expose.jar"}
 repl {:eval (use 'play.core)})


(deftask build
  "build my project"
  []
  (comp
   (aot :all true)
   (uber)
   (jar)
   (target)))

(require '[play.counting])

(deftask run []
  (with-pre-wrap fileset
    (play.counting/-main)
    fileset))
