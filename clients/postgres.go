package main

import (
  "context"
  "database/sql"
  "encoding/csv"
  "fmt"
  "io"
  "os"
  "strings"

  _ "github.com/lib/pq"
)

var (
  ctx context.Context
  db *sql.DB
)

/**
  ************ POSTGRES DB INSERTER ************
  Read a csv file gathered from rpscrape and insert into lingfield table (probably change that soon to be all courses/regions)
  On error just panic / don't try and salvage
  Rollsback error'd transactions though
*/
func main() {
  // establish db connection
  cs := "postgresql://localhost/horse_racing?sslmode=disable"
  conn, err := sql.Open("postgres", cs)

  if err != nil {
    panic(err)
  }

  // read horse_racing data csv
  // d, err := os.ReadFile("./lingfield_2023_flat_no_title.csv")
  d, err := os.ReadFile("./one_row.csv")
  if err != nil {
    panic(err)
  }
  
  r := csv.NewReader(strings.NewReader(string(d)))

  for {
    // read data row
    record, err := r.Read()
    if err == io.EOF {
      break;
    }
    if err != nil {
      panic(err)
    }

    // create transaction
    tx, err := db.BeginTx(ctx, &sql.TxOptions{})
    if err != nil {
      panic(err):
    }

    // do once?
    t := strings.Repeat("?,", len(record))
    v := strings.TrimRight(t, ",")
    fmt.Println(v)
 
    p := make([]interface{}, len(record))
    for i, x := range record {
      p[i] = x
    }

    // exec
    _, err = tx.Exec("INSERT INTO lingfield VALUES(" + v + ");", p...)
    if err != nil {
      // rollback if it fs up, but no big deal currently, can just truncate table (16.NOV.2023)
      _ = tx.Rollback()
      panic(err)
    }

    // commit
    if err := tx.Commit(); err != nil {
      panic(err)
    }

    fmt.Println(record[0])
    // repeat
  }

	rows, err := conn.Query("SELECT count(*) from lingfield;")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var version string
		rows.Scan(&version)
		fmt.Println(version)
	}

	rows.Close()


  conn.Close()
}
