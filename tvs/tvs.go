package tvs

import (
	"fmt"
	"net/http"
	"rus-sharafiev/kodi/tvs/queries"

	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	result := queries.FindAll()
	fmt.Fprint(w, result)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	result := queries.FindOne(id)
	fmt.Fprint(w, result)
}
