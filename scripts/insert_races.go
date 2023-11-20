package main

import (
  "fmt"
  "strconv"
  "database/sql"
  "encoding/csv"
  "io"
  "os"
  "strings"

  "github.com/urfave/cli/v2"
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
        var commit bool
        var filepath string

        app := &cli.App{
                Name: "Result Insert",
                Usage: "Extracts data from a csv from rpscrape, transform it into something insertable into local PostgresDB",
                Action: func(*cli.Context) error {
                        
                        return run(commit, filepath)                
                },
                Flags: 
                        []cli.Flag{
                                &cli.BoolFlag{
                                        Name: "commit",
                                        Aliases: []string{"c"},
                                        Value: false,
                                        Usage: "insert data into real table rather than temp",
                                        Destination: &commit,
                                },
                                
                                &cli.StringFlag{
                                        Name: "filepath",
                                        Aliases: []string{"f"},
                                        Usage: "path to csv horse-racing file to extract",
                                        Destination: &filepath,
                                },
                        },
                }

        if err := app.Run(os.Args); err != nil {
                panic(err)
        }
}

func run(commit bool, filepath string) error {

        // Open DB connection
        cs := "postgresql://localhost/horse_racing?sslmode=disable"
        db, err := sql.Open("postgres", cs)
        if err != nil {
                panic(err)
        }

        // Read CSV 
        d, err := os.ReadFile(filepath)
        if err != nil {
                panic(err)
        }
  
        r := csv.NewReader(strings.NewReader(string(d)))
        rn := 1

        tn := "temp"
        if commit {
                tn = "lingfield"
        }

        // Process CSV
        for {
                record, err := r.Read()
                if err == io.EOF {
                        break;
                }
                if err != nil {
                        return err
                }

                tx, err := db.Begin()
                if err != nil {
                        return err
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

                // TODO only need to do this once
                v := strings.TrimRight(sb.String(), ",") // remove trailing comma
                st := fmt.Sprintf("INSERT INTO %s VALUES(%s);", tn, v)
                _, err = tx.Exec(st, p...)
                if err != nil {
                        _ = tx.Rollback()
                        fmt.Println("* row: ", rn)
                        return err
                }

                if err := tx.Commit(); err != nil {
                        return err      
                }
                rn++
        }

        // Finished
        fmt.Println(fmt.Sprintf("%s rows processed %d", filepath, rn))

        return db.Close()
}

type root struct {}
type data struct {}

// TODO Make a different type for position, key pair? 
// FI - int
// PU - NULL
// UR - NULL
// DSQ - NULL
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
