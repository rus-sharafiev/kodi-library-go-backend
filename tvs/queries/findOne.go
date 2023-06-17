package queries

import (
	"fmt"
	"log"
	"rus-sharafiev/kodi/db"
	"strings"
)

func FindOne(id string) string {

	var query strings.Builder

	query.WriteString(`
		SELECT IFNULL(
			JSON_OBJECT(
				'art', a.arr,
				'id', tv.idShow,
				'title', tv.c00,
				'description', tv.c01,
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
				'totalCount', tv.totalCount,

				'seasons', s.arr
			)
		, 'null') result
		
		FROM tvshow_view tv

		LEFT JOIN (
			SELECT media_id, JSON_OBJECTAGG(type, url) arr
			FROM art
			GROUP BY art.media_id
		) a
		ON tv.idShow = a.media_id

		LEFT JOIN (
			SELECT idShow, JSON_ARRAYAGG(
				JSON_OBJECT(
					'idSeason', sv.idSeason,
					'idShow', sv.idShow,
					'season', sv.season,
					'name', sv.name,
					'userrating', sv.userrating,
					'strPath', sv.strPath,
					'showTitle', sv.showTitle,
					'plot', sv.plot,
					'premiered', sv.premiered,
					'genre', sv.genre,
					'studio', sv.studio,
					'mpaa', sv.mpaa,
					'episodes', sv.episodes,
					'playCount', sv.playCount,
					'aired', sv.aired,
					
					'art', a.arr,
					'episodes', e.arr
				)
			) arr
			FROM season_view sv

			LEFT JOIN (
				SELECT idSeason, JSON_ARRAYAGG(
					JSON_OBJECT(
						'idEpisode', idEpisode,
						'idFile', idFile,
						'title', c00, 
						'description', c01, 
						'idShow', idShow,
						'userrating', userrating,
						'idSeason', idSeason,
						'strFileName', strFileName,
						'strPath', strPath,
						'playCount', playCount,
						'lastPlayed', lastPlayed,
						'dateAdded', dateAdded,
						'strTitle', strTitle,
						'genre', genre,
						'studio', studio,
						'premiered', premiered,
						'mpaa', mpaa,
						'resumeTimeInSeconds', resumeTimeInSeconds,
						'totalTimeInSeconds', totalTimeInSeconds,
						'playerState', playerState,
						'rating', rating,
						'votes', votes,
						'rating_type', rating_type,
						'uniqueid_value', uniqueid_value,
						'uniqueid_type', uniqueid_type
					)
				) arr
				FROM episode_view
				GROUP BY episode_view.idSeason
			) e
			ON sv.idSeason = e.idSeason

			LEFT JOIN (
				SELECT media_id, JSON_OBJECTAGG(type, url) arr
				FROM art
				GROUP BY art.media_id
			) a
			ON sv.idSeason = a.media_id

			GROUP BY sv.idShow
		) s
		ON tv.idShow = s.idShow
	`)

	fmt.Fprintf(&query, "WHERE tv.idShow = %s;", id)

	db := db.Connect()
	defer db.Close()

	var result string
	if err := db.QueryRow(query.String()).Scan(&result); err != nil {
		log.Fatal(err)
	}

	return result
}
