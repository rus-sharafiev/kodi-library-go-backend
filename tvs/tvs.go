package tvs

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
			'id', tv.idShow,
			'title', tv.c00,
			'review', tv.c01,
			'premiered', tv.c05,
			'genre', tv.c08,
			'originalTitle', tv.c09,
			'pgRating', tv.c13,
			'studio', tv.c14,
			'youtube', tv.c16,
			'rating', tv.rating,
			'votes', tv.votes,
			'ratingType', tv.rating_type,

			'watchedcount', tv.watchedcount,
			'totalSeasons', tv.totalSeasons,
			'lastPlayed', tv.lastPlayed,
			'dateAdded', tv.dateAdded,
			'totalCount', tv.totalCount
			`)

	if len(id) > 0 {
		query.WriteString("), 'null') result")
	} else {
		query.WriteString(")), 'null') result")
	}

	query.WriteString(`	
		FROM tvshow_view tv
		LEFT JOIN (
			SELECT media_id, JSON_OBJECTAGG(type, url) arr
			FROM art
			group by art.media_id
		) a
		ON tv.idShow = a.media_id
	`)

	if len(id) > 0 {
		fmt.Fprintf(&query, "WHERE tv.idShow = %s;", id)
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
