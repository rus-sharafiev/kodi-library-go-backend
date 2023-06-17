package movies

import (
	"fmt"
	"log"
	"net/http"
	"rus-sharafiev/kodi/db"
)

func selectAll() string {
	db := db.Connect()
	defer db.Close()

	query := `
		SELECT JSON_ARRAYAGG(
			JSON_OBJECT(
				'art', a.arr,
				'idMovie', m.idMovie,
				'title', m.c00,
				'review', m.c01,
				'subtitle', m.c03,
				'writer', m.c06,
				'pg_rating', m.c12,
				'genre', m.c14,
				'director', m.c15,
				'original_title', m.c16,
				'studio', m.c18,
				'youtube', m.c19,
				'country', m.c21,
				'premiered', m.premiered,
				'playCount', m.playCount,
				'lastPlayed', m.lastPlayed,
				'dateAdded', m.dateAdded,
				'rating', m.rating,
				'votes', m.votes
			)) result
		FROM movie_view m
		LEFT JOIN (
			SELECT media_id, JSON_OBJECTAGG(type, url) arr
			FROM art
			group by art.media_id
		) a
		ON m.idMovie = a.media_id;
	`
	var result string
	if err := db.QueryRow(query).Scan(&result); err != nil {
		log.Fatal(err)
	}

	return result
}

func Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		result := selectAll()
		fmt.Fprint(w, result)
	})
}
