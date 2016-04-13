(ns play.core)

(def hello-world "++++++++ [>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-] >>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")

(defrecord VM [text pc addr data])

(defn match [text index dir open-brackets]
  (let [open-br (case (get text index)
                  \[ (inc open-brackets)
                  \] (dec open-brackets)
                  open-brackets)]
    (if (zero? open-br)
      index
      (recur text (dir index) dir open-br))))

(defn step [{:keys [text pc addr data]}]
  (let [next (->VM text (inc pc) addr data)]
    (case (get text pc)
      \+ (update-in next [:data addr] (fnil inc 0))
      \- (update-in next [:data addr] (fnil dec 0))
      \> (update next :addr inc)
      \< (update next :addr dec)
      \. (do (print (char (get data addr)))
             next)
      \, next
      \[ (if (zero? (get data addr 0))
           (assoc next :pc (inc (match text pc inc 0)))
           next)
      \] (assoc next :pc (match text pc dec 0))
      next)))

(defn running [prog]
  (< (:pc prog)
     (.length (:text prog))))

(defn eval-bf
  "evaluate a BF program"
  [source]
  (loop [prog (->VM source 0 0 {})]
    (if (running prog)
      (recur (step prog)))))
