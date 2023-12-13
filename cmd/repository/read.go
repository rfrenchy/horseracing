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
        st := "SELECT csv FROM racingpost WHERE course_id = $1 AND year = $2"

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
