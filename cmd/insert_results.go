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
        var commit bool // TODO change to string flag  (raw | lingfield | result)
        var filepath string

        app := &cli.App{
                Name: "Result Insert",
                Usage: "Extracts and processes horse-racing results CSV, inserts into PostresDB",
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

type runargs struct {
        filepath string
        table string
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

        // Skip header line
        _, err = r.Read()
        if err != nil {
                return err;
        }

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

                // TODO only need to do this once, 
                v := strings.TrimRight(sb.String(), ",") // remove trailing comma
                tn := "temp"
                if commit {
                        tn = "lingfield" // TODO get proper table name
                }
                st := fmt.Sprintf("INSERT INTO %s VALUES(%s);", tn, v)
                _, err = tx.Exec(st, p...)
                if err != nil {
                        _ = tx.Rollback()
                        fmt.Println("* error row:", rn)
                        e := err
                        f, err := os.OpenFile("err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	                if err != nil {
		                return err
	                }
                        defer f.Close()

                        if _, err = f.WriteString(fmt.Sprintf("%d:%s\n%v\n", rn, e, record)); err != nil {
                                return err
                        }
                        rn++
                        continue
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

type root struct{}

func (p *root) position(s string) string {
        switch s {
                default:    return s
                case "PU":  return strconv.Itoa(-1) // (Pulled up i.e. injury/issue)
                case "UR":  return strconv.Itoa(-2) // (Unseated Rider)
                case "DSQ": return strconv.Itoa(-3) // (Disqualified)
                case "SU":  return strconv.Itoa(-4) // 
                case "F":   return strconv.Itoa(-5) // (Fell)
                case "RR":  return strconv.Itoa(-6) // (Refused to Race)
                case "BD":  return strconv.Itoa(-7) // (Brought down)
        }
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
