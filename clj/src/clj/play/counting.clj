(ns play.counting)

(def digits
  (set (map #(char (+ 48 %))
            (range 10))))

(defn left [digits num]
  (clojure.set/difference digits (set (str num))))

(defn solve [n]
  (if (= n 0) "INSOMNIA"
      (loop [i 1
             needed digits]
        (let [num (* i n)
              needed (left needed num)]
          (if (empty? needed) num
              (recur (inc i) needed))))))

(defn printer [[index result]]
  (println (str "Case #" (inc index) ": " result)))

(defn solve-string [solver str]
  (->> str
      clojure.string/split-lines
      rest
      (map #(-> (read-string %)
                solve))
      (map vector (range))
      (map printer)))

(defn -main []
  (doall (solve-string solve (slurp *in*))))
