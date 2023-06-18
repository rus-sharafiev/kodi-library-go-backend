package main

import (
	"fmt"
	"log"
	"net/http"

	"rus-sharafiev/kodi/movies"
	"rus-sharafiev/kodi/tvs"
	"rus-sharafiev/kodi/web"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	// API
	api := router.PathPrefix("/api").Subrouter()

	// Movies
	api.HandleFunc("/movies", movies.GetAll).Methods("GET")
	api.HandleFunc("/movies/{id}", movies.GetOne).Methods("GET")

	// Tvs
	api.HandleFunc("/tvs", tvs.GetAll).Methods("GET")
	api.HandleFunc("/tvs/{id}", tvs.GetOne).Methods("GET")

	// Web server
	router.PathPrefix("/").Handler(web.Server{StaticPath: "build", IndexPath: "index.html"})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	// Start server
	fmt.Printf("\n\x1b[2mHTTP server is running on http://localhost:8088/\n\n\x1b[0m")
	log.Fatal(http.ListenAndServe(":8088", handler))
}
