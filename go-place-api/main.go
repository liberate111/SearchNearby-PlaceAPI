package main

import (
	"fmt"
	"maps/place"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
	})

	http.HandleFunc("/", index)
	http.HandleFunc("/xyz", place.FindValue)
	http.HandleFunc("/nearby", place.SearchNearby)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8090", c.Handler(http.DefaultServeMux))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Google Place API")
}

// example
// http://localhost:8080/
// http://localhost:8080/xyz
// http://localhost:8080/nearby?name=<some place>
