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

	movies.Controller(api)
	tvs.Controller(api)

	// Web server
	router.PathPrefix("/").HandlerFunc(web.Server)

	// CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://10.10.10.10:8000", "http://localhost:8000", "http://localhost:8088"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	// Start server
	fmt.Printf("\n\x1b[2mHTTP server is running on http://localhost:8088/\n\n\x1b[0m")
	log.Fatal(http.ListenAndServe(":8088", handler))
}
