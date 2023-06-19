package tvs

import (
	"fmt"
	"net/http"
	"rus-sharafiev/kodi/tvs/queries"

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
	tvs := api.PathPrefix("/tvs").Subrouter()

	tvs.HandleFunc("", findAll).Methods("GET")
	tvs.HandleFunc("{id}/", findOne).Methods("GET")
}
