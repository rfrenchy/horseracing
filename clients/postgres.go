package main

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)


func main() {
  cs := "postgresql://localhost/horse_racing?sslmode=disable"

  conn, err := sql.Open("postgres", cs)

  if err != nil {
    panic(err)
  }


  // CONNECTION OPEN
  fmt.Println("")

  // READ FILE
  fmt.Println("")


  // INSERT DATA

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
