package main

import (
	"fmt"
	"log"
	"net/http"
	"rus-sharafiev/kodi/movies"
	"rus-sharafiev/kodi/web"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Web server
	router.Handle("/", web.Server())

	// API
	router.Handle("/api/movies", movies.Get()).Methods("GET")

	fmt.Printf("\n\x1b[2mHTTP server is running on http://localhost:8088/\n\n\x1b[0m")
	log.Fatal(http.ListenAndServe(":8088", router))
}
