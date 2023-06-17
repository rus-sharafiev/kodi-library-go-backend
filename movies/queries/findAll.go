package queries

import (
	"log"
	"rus-sharafiev/kodi/db"
	"strings"
)

func FindAll() string {

	var query strings.Builder

	query.WriteString(`
		SELECT JSON_ARRAYAGG(
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
			)
		) result
		FROM movie_view m

		LEFT JOIN (
			SELECT media_id, JSON_OBJECTAGG(type, url) arr
			FROM art
			GROUP BY art.media_id
		) a
		ON m.idMovie = a.media_id;
	`)

	db := db.Connect()
	defer db.Close()

	var result string
	if err := db.QueryRow(query.String()).Scan(&result); err != nil {
		log.Fatal(err)
	}

	return result
}
