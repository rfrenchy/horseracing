package main

import (
	"database/sql"

	"github.com/gocarina/gocsv"
	_ "github.com/lib/pq"
)

// Reader contains query's on horse racing
type Read struct {
	db *sql.DB
}

// RacingPostRecords returns all records for a given course for the given year
func (read *Read) RacingPost(cid int, year int) ([]*RacingPostRecord, error) {
	// Query statement
	st := `	SELECT csv 
		FROM racingpost 
		WHERE course_id = $1 AND year = $2`

	// Read CSV
	var csv []byte
	if err := read.db.QueryRow(st, cid, year).Scan(&csv); err != nil {
		return nil, err
	}

	// Unmarshall to records
	var records []*RacingPostRecord
	if err := gocsv.UnmarshalBytes(csv, &records); err != nil {
		return nil, err
	}

	// Return records
	return records, nil
}

// TODO, change to parametrised entity i.e. Trainer | Owner | Horse | Jockey | ...

// TrainerPositionRecords
// func (read *Read) TrainerPositionRecords() error {
// 	// Query statement
// 	st: = "	SELECT count(ru.\"position\"), ru.\"position\"
// 		FROM course c
// 		INNER JOIN race r on r.course_id = c.course_id
// 		INNER JOIN runner ru on ru.race_id = r.race_id
// 		WHERE ru.trainer_id = 9714 and c.course_id = 7
// 		GROUP BY ru.\"position\"
// 		ORDER BY \"count\" DESC"
//
// 	return nil
// }
//
// // TrainerWinRates gets how often the trainer of a horse wins ordered from best to worst
// func (read *Read) TrainerWinRates() error {
// 	// Query statement
// 	st := "	SELECT COUNT(r.trainer_id) as wins, r.trainer_id, e2.races, (COUNT(r.trainer_id) / cast(e2.races as decimal)) as win_rate
// 		FROM runner r
// 		INNER JOIN (SELECT COUNT(r.race_id) as races, r.trainer_id from runner r group by r.trainer_id) e2 on e2.trainer_id = r.trainer_id
// 		WHERE \"position\" = 1
// 		GROUP BY r.trainer_id, e2.races
// 		ORDER BY win_rate DESC"
//
// 	return nil
// }
//
// // HorseWinRates
// func (read *Read) HorseWinRates() error {
// 	st := "	SELECT COUNT(r.horse_id) as wins, r.horse_id, e2.races, (COUNT(r.horse_id) / cast(e2.races as decimal)) as win_rate
// 		FROM runner r
// 		INNER JOIN (SELECT COUNT(r.race_id) as races, r.horse_id from runner r group by r.horse_id) e2 on e2.horse_id = r.horse_id
// 		WHERE "position" = 1
// 		GROUP BY r.horse_id, e2.races
// 		ORDER BY races DESC, win_rate DESC"
//
// 	return nil
// }
