(ns backend.util)

(defn dist [c1 c2]
  (->> (map - c1 c2) (map #(* % %)) (reduce +)))

(defn item-dist [c {lat :lat lon :lon}]
  (dist c [lat lon]))

(defn sort-closest [my-lat my-lon item-list]
  (sort-by (partial item-dist [my-lat my-lon]) item-list))

(defn find-closest [my-lat my-lon item-list]
  (apply min-key (partial item-dist [my-lat my-lon]) item-list))
