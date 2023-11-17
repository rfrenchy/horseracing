package main

import (
  "fmt"
  "strconv"
  "database/sql"
  "encoding/csv"
  "io"
  "os"
  "strings"

  _ "github.com/lib/pq"
)

/**
  ************ POSTGRES DB INSERTER ************
  
  - Read a csv file gathered from rpscrape and insert into lingfield table
    * lingfield name will be changed soon to course/region

  - On error just panic / don't try and salvage
    * rollsback error'd transactions though

  **********************************************
*/
func main() {
  // Establish DB connection
  cs := "postgresql://localhost/horse_racing?sslmode=disable"
  db, err := sql.Open("postgres", cs)

  if err != nil {
    panic(err)
  }

  // Read horse_racing data CSV
  d, err := os.ReadFile("./lingfield_2023_flat_no_title.csv")
  if err != nil {
    panic(err)
  }
  
  r := csv.NewReader(strings.NewReader(string(d)))
  ri := 0
  for {
    fmt.Println("Processing row", ri)
    // Read CSV data row
    record, err := r.Read()
    if err == io.EOF {
      break;
    }
    if err != nil {
      panic(err)
    }

    // Create a Transaction
    tx, err := db.Begin()
    if err != nil {
      panic(err)
    }

    var sb strings.Builder
    p := make([]interface{}, len(record)) // can optimise by only working out len(record) once
    for i, x := range record {
      // create param string for SQL i.e. $1, $2, $3...
      sb.WriteString("$" + strconv.Itoa(i + 1) + ",")      

      if i == 31 && (x == "" || x == "–") {
        p[i] = 0 // convert prize column data to 0 if there is no prize
      } else if i == 33 && x == "–" {
        p[i] = 0 // convert RPR to 0 if "-" assessment i.e. no assessment score
      } else if i == 17 && x == "PU" {
        p[i] = -1 // use -1 to resemble Pulled Up (PU) finish (deliberate stop by jockey)
      } else if i == 19 && x == "-" {
        p[i] = 0
      } else if i == 20 && x == "-" {
        p[i] = 0
      } else if i == 26 && x == "-" {
        p[i] = nil
      } else if i == 27 && x == "-" {
        p[i] = nil
      } else {
        p[i] = x
      }
    }

    v := strings.TrimRight(sb.String(), ",") // remove trailing comma

    // Execute INSERT statement into lingfield table
    st := "INSERT INTO lingfield VALUES(" + v + ");"

    _, err = tx.Exec(st, p...)
    if err != nil {
      _ = tx.Rollback()
      panic(err)
    }

    if err := tx.Commit(); err != nil {
      panic(err)
    }

    ri++;
    // Repeat for rest of CSV
  }

  db.Close()
}
