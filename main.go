package main

import (
	"fmt"
	"log"
	"net/http"
	"rus-sharafiev/kodi/movies"
	"rus-sharafiev/kodi/tvs"
	"rus-sharafiev/kodi/web"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// API
	router.Handle("/api/movies", movies.Get()).Methods("GET")
	router.Handle("/api/movies/{id}", movies.Get()).Methods("GET")

	router.Handle("/api/tvs", tvs.Get()).Methods("GET")
	router.Handle("/api/tvs/{id}", tvs.Get()).Methods("GET")

	// Web server
	router.PathPrefix("/").Handler(web.Server{StaticPath: "build", IndexPath: "index.html"})

	fmt.Printf("\n\x1b[2mHTTP server is running on http://localhost:8088/\n\n\x1b[0m")
	log.Fatal(http.ListenAndServe(":8088", router))
}
