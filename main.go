package main

import "net/http"

func main() {
	http.HandleFunc("/scatter", ScatterHandler)
	http.HandleFunc("/square", SquareHandler)
	http.HandleFunc("/grid", GridHandler)
	http.HandleFunc("/image", ImageHandler)
	http.HandleFunc("/hist", HistogramHandler)

	http.ListenAndServe(":8998", nil)
}
