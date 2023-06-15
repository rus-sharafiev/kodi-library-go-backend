package main

import (
	"fmt"
	"log"
	"net/http"
	"rus-sharafiev/kodi/fswr"
	"rus-sharafiev/kodi/movies"
)

func main() {
	// App server
	http.Handle("/", http.StripPrefix("/", fswr.FileServerWithRedirect(http.Dir("build/"))))

	// API
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		result := movies.GetAll()
		fmt.Fprint(w, result)
	})

	fmt.Printf("\n\x1b[2mHTTP server is running on http://localhost:8088/\n \x1b[0m ")

	log.Fatal(http.ListenAndServe(":8088", nil))
}
