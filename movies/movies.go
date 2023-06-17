package movies

import (
	"fmt"
	"log"
	"net/http"
	"rus-sharafiev/kodi/db"
	"strings"

	"github.com/gorilla/mux"
)

func queryMovies(id string) string {
	db := db.Connect()
	defer db.Close()

	var query strings.Builder

	if len(id) > 0 {
		query.WriteString("SELECT IFNULL(")
	} else {
		query.WriteString("SELECT IFNULL(JSON_ARRAYAGG(")
	}

	query.WriteString(`
	JSON_OBJECT(
		'art', a.arr,
		'id', m.idMovie,
		'title', m.c00,
		'description', m.c01,
		'subtitle', m.c03,
		'writer', m.c06,
		'pgRating', m.c12,
		'genre', m.c14,
		'director', m.c15,
		'originalTitle', m.c16,
		'studio', m.c18,
		'youtube', m.c19,
		'country', m.c21,
		'premiered', m.premiered,
		'rating', m.rating,
		'votes', m.votes,
		'ratingType', m.rating_type,

		'playCount', m.playCount,
		'lastPlayed', m.lastPlayed,
		'dateAdded', m.dateAdded,
		'resumeTimeInSeconds', m.resumeTimeInSeconds,
		'totalTimeInSeconds', m.totalTimeInSeconds
	`)

	if len(id) > 0 {
		query.WriteString("), 'null') result")
	} else {
		query.WriteString(")), 'null') result")
	}

	query.WriteString(`	
		FROM movie_view m
		LEFT JOIN (
			SELECT media_id, JSON_OBJECTAGG(type, url) arr
			FROM art
			GROUP BY art.media_id
		) a
		ON m.idMovie = a.media_id
	`)

	if len(id) > 0 {
		fmt.Fprintf(&query, "WHERE m.idMovie = %s;", id)
	} else {
		query.WriteString(";")
	}

	var result string
	if err := db.QueryRow(query.String()).Scan(&result); err != nil {
		log.Fatal(err)
	}

	return result
}

func Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		id := mux.Vars(r)["id"]
		result := queryMovies(id)
		fmt.Fprint(w, result)
	})
}
