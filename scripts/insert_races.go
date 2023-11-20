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
  ************ HORSE RACING POSTGRES INSERTER ************
  
  - Read a CSV file gathered from rpscrape and insert into lingfield table
    * lingfield name will be changed soon to course/region

  - On error just panic / don't try and salvage
    * rollsback error'd transactions though

  - file as input arg?

  **********************************************
*/
func main() {

        // Establish DB connection
        cs := "postgresql://localhost/horse_racing?sslmode=disable"
        db, err := sql.Open("postgres", cs)
        if err != nil {
                panic(err)
        }

        // Read CSV into var
//        path := "./lingfield_2023_flat_no_title.csv"
        path := "./temp/2021.csv"
        d, err := os.ReadFile(path)
        if err != nil {
                panic(err)
        }
  
        r := csv.NewReader(strings.NewReader(string(d)))
        rn := 1

        fmt.Println("Processing CSV")
        for {
                fmt.Println("* row: ", rn)
                
                record, err := r.Read()
                if err == io.EOF {
                        break;
                }
                if err != nil {
                        panic(err)
                }

                tx, err := db.Begin()
                if err != nil {
                        panic(err)
                }

                var sb strings.Builder
                p := make([]interface{}, len(record))

                pr := &root{}

                for i, x := range record {
                        // $1, $2, $3... etc
                        sb.WriteString("$" + strconv.Itoa(i + 1) + ",")      
                       
                        switch i {
                                default: p[i] = x
                                case 17: p[i] = pr.position(x)
                                case 19: p[i] = pr.ovrbtn(x) 
                                case 20: p[i] = pr.btn(x) 
                                case 26: p[i] = pr.time(x)
                                case 27: p[i] = pr.seconds(x) 
                                case 31: p[i] = pr.prize(x)
                                case 33: p[i] = pr.rpr(x) 
                        }
                }

                v := strings.TrimRight(sb.String(), ",") // remove trailing comma

                //st := "INSERT INTO lingfield VALUES(" + v + ");"
                st := "INSERT INTO temp VALUES(" + v + ");"

                _, err = tx.Exec(st, p...)
                if err != nil {
                        _ = tx.Rollback()
                        panic(err)
                }

                if err := tx.Commit(); err != nil {
                        panic(err)
                }

                rn++
        }

        db.Close()
}


type root struct {}
type data struct {}

// TODO Make a different type for finishes, key pair? 
func (p *root) position(s string) string {
        if s == "PU" { // pulled up
                return strconv.Itoa(-1)
        } else if s == "UR" {
                return strconv.Itoa(-2)
        } else if s == "DSQ" {
                return strconv.Itoa(-3)
        }
        return s
}

func (p *root) ovrbtn(s string) string {
        if s == "-" { 
                return strconv.Itoa(0)
        }
        return s
}

func (p *root) btn(s string) string {
        if s == "-" { 
                return strconv.Itoa(0)
        }
        return s
}

func (p *root) time(s string) *string {
        if s == "-" { 
                return nil
        }
        return &s
}

func (p *root) seconds(s string) *string {
        if s == "-" { 
                return nil
        }
        return &s
}

func (p *root) prize(s string) string {
        if s == "" || s == "–" {
                return strconv.Itoa(0)
        } 
        return s
}

func (p *root) rpr(s string) string {
        if s == "" || s == "–" {
                return strconv.Itoa(0)
        } 
        return s
}
