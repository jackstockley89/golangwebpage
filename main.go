package main

import (
	"fmt"
	"net/http"

	"github.com/jackstockley89/golangwebpage/lib"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		lib.HomeHandler(w, r)
	})

	http.HandleFunc("/activities", func(w http.ResponseWriter, r *http.Request) {
		accept := r.Header.Get("Accept")
		wantJson := accept == "application/json"
		lib.ActivitiesHandler(w, r, wantJson)
	})

	http.HandleFunc("/activities/", lib.ActivityDetailHandler)

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
