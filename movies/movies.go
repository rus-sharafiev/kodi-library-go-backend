package movies

import (
	"fmt"
	"net/http"
	"rus-sharafiev/kodi/movies/queries"

	"github.com/gorilla/mux"
)

func findAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, queries.FindAll())
}

func findOne(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	fmt.Fprint(w, queries.FindOne(id))
}

func Controller(api *mux.Router) {
	movies := api.PathPrefix("/movies").Subrouter()

	movies.HandleFunc("", findAll).Methods("GET")
	movies.HandleFunc("/{id}", findOne).Methods("GET")
}
