package db

import (
  "fmt"
  "strconv"
  "database/sql"
  "encoding/csv"
  "io"

  _ "github.com/lib/pq"
)

type Write struct {
// db connecion
}

type root struct{}

func RacingPost(f *File) error {
	records := []*RacingPostRecord{}
	if err := gocsv.UnmarshalFile(csv, &records); err != nil {
		return err
	}

        // start at 1, skip header line
        for i := 1; i < len(records); i++ {
                // EXEC
                st := fmt.Sprintf("INSERT INTO %s VALUES(%s);", tn, v)
                _, err = tx.Exec(st, p...)
                if err != nil {
                        _ = tx.Rollback()

                        log.Error().Err(err).Int("row", i).
                                Interface("params", p).
                                Msg("SKIPPING ROW")

                        continue
                }

                // COMMIT
                if err := tx.Commit(); err != nil {
                        return err
                }
        }
}

func (w *Write) Horse() error {

}

func (w *Write) Owner() error {

}

func (w *Write) Jockey() error {

}

func (w *Write) Race() error {

}

func (w *Write) Runner() error {

}

func (p *root) num(s string) string {
        if s == "" {
                return strconv.Itoa(0) // blank?
        }
        return s
}

// Abbreviations help - https://help.racingpost.com/hc/en-us/articles/115001699689-Abbreviations-on-the-racecard
func (p *root) position(s string) string {
        switch s {
                default:     return s
                case "PU":   return strconv.Itoa(-1) // (Pulled up i.e. injury/issue)
                case "UR":   return strconv.Itoa(-2) // (Unseated Rider)
                case "DSQ":  return strconv.Itoa(-3) // (Disqualified)
                case "SU":   return strconv.Itoa(-4) // ?
                case "F":    return strconv.Itoa(-5) // (Fell)
                case "RR":   return strconv.Itoa(-6) // (Refused to Race)
                case "BD":   return strconv.Itoa(-7) // (Brought down)
                case "LFT":  return strconv.Itoa(-8) // ?
                case "RO":   return strconv.Itoa(-9) // (Refused to Race?)
        }
}

func (p *root) draw(s string) string {
        if s == "" {
                return strconv.Itoa(0) // blank?
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

               // for i, x := range record {
               //         // $1, $2, $3... etc
               //         switch i {
               //                 default: p[i] = x
               //                 case 16: p[i] = pr.num(x)
               //                 case 17: p[i] = pr.position(x)
               //                 case 18: p[i] = pr.draw(x)
               //                 case 19: p[i] = pr.ovrbtn(x)
               //                 case 20: p[i] = pr.btn(x)
               //                 case 26: p[i] = pr.time(x)
               //                 case 27: p[i] = pr.seconds(x)
               //                 case 31: p[i] = pr.prize(x)
               //                 case 33: p[i] = pr.rpr(x)
                        }
