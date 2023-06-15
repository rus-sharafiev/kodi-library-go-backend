package movies

import (
	"log"
	"rus-sharafiev/kodi/db"
)

func GetAll() string {
	db := db.Connect()
	defer db.Close()

	query := `
		SELECT JSON_ARRAYAGG(JSON_OBJECT('title', m.c00, 'subtitle', m.c03, 'arts', a.arr)) result 
		FROM movie_view m
		LEFT JOIN (
			SELECT media_id, JSON_ARRAYAGG(JSON_OBJECT(type, url)) arr
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
