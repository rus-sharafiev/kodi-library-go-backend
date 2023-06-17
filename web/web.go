package web

import (
	"net/http"
	"os"
	"path/filepath"
)

func Server() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path, err := filepath.Abs(r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		path = filepath.Join("build", path)

		_, err = os.Stat(path)
		if os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join("build", "index.html"))
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.FileServer(http.Dir("build")).ServeHTTP(w, r)
	})
}
