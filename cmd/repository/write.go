package main

import (
  "fmt"
  "strconv"
  "errors"
  "database/sql"

  _ "github.com/lib/pq"
)

type Write struct {
        db *sql.DB
}

type root struct{}

func (w *Write) Add(r *RacingPostRecord) error {
        err_own := w.Owner(r)
        err_jky := w.Jockey(r)
        err_trn := w.Trainer(r)
        err_hrs := w.Horse(r)
        err_rce := w.Race(r)
        err_run := w.Runner(r)

        return errors.Join(err_own, err_jky, err_trn, err_hrs, err_rce, err_run)
}

// RacingpostCSV persists mined Racingpost CSV
func (w *Write) Racingpost(cid int, year int, csv *[]byte) error {
        tx, err := w.db.Begin()
        if err != nil {
                return err
        }

        st := fmt.Sprintf("INSERT INTO racingpost VALUES($1, $2, $3, $4) ON CONFLICT DO NOTHING;")

        _, err = tx.Exec(st, cid, year, csv, false)
        if err != nil {
                _ = tx.Rollback()
                return err
        }

        if err := tx.Commit(); err != nil {
                return err
        }

        return nil
}

func (w *Write) Owner(r *RacingPostRecord) error {
        tx, err := w.db.Begin()
        if err != nil {
                return err
        }

        st := fmt.Sprintf("INSERT INTO owner VALUES($1, $2, $3) ON CONFLICT DO NOTHING;")

        _, err = tx.Exec(st, r.OwnerID, r.OwnerName, r.SilkURL)
        if err != nil {
                _ = tx.Rollback()
                return err
        }

        if err := tx.Commit(); err != nil {
                return err
        }

        return nil
}

func (w *Write) Jockey(r *RacingPostRecord) error {
        tx, err := w.db.Begin()
        if err != nil {
                return err
        }

        st := fmt.Sprintf("INSERT INTO jockey VALUES($1, $2) ON CONFLICT DO NOTHING;")

        _, err = tx.Exec(st, r.JockeyID, r.JockeyName)
        if err != nil {
                _ = tx.Rollback()
                return err
        }

        if err := tx.Commit(); err != nil {
                return err
        }

        return nil
}

func (w *Write) Trainer(r *RacingPostRecord) error {
        tx, err := w.db.Begin()
        if err != nil {
                return err
        }

        st := fmt.Sprintf("INSERT INTO trainer VALUES($1, $2) ON CONFLICT DO NOTHING;")

        _, err = tx.Exec(st, r.TrainerID, r.TrainerName)
        if err != nil {
                _ = tx.Rollback()
                return err
        }

        if err := tx.Commit(); err != nil {
                return err
        }

        return nil
}

func (w *Write) Horse(r *RacingPostRecord) error {
        tx, err := w.db.Begin()
        if err != nil {
                return err
        }

        st := fmt.Sprintf("INSERT INTO horse VALUES($1, $2, $3) ON CONFLICT DO NOTHING;")

        _, err = tx.Exec(st, r.HorseID, r.HorseName, r.HorseSex)
        if err != nil {
                _ = tx.Rollback()
                return err
        }

        if err := tx.Commit(); err != nil {
                return err
        }

        return nil
}

func (w *Write) Race(r *RacingPostRecord) error {
        tx, err := w.db.Begin()
        if err != nil {
                return err
        }

        st := fmt.Sprintf("INSERT INTO race VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) ON CONFLICT DO NOTHING;")

        _, err = tx.Exec(st,
                r.RaceID,
                r.RaceName,
                r.RaceDate,
                r.CourseID,
                r.RaceOfftime,
                r.RaceType,
                r.RaceClass,
                r.RacePattern,
                r.RatingbandRestrictions,
                r.AgebandRestriction,
                r.SexRestriction,
                r.Distance,
                r.Going,
                r.Surface,
                r.Ran)

        if err != nil {
                _ = tx.Rollback()
                return err
        }

        if err := tx.Commit(); err != nil {
                return err
        }

        return nil
}

func (w *Write) Runner(r *RacingPostRecord) error {
        tx, err := w.db.Begin()
        if err != nil {
                return err
        }

        st := fmt.Sprintf("INSERT INTO runner VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20) ON CONFLICT DO NOTHING;")

        root := &root{}

        _, err = tx.Exec(st,
                r.HorseID,
                r.RaceID,
                root.num(r.RacecardNumber),
                root.position(r.FinishedPosition),
                root.draw(r.Draw),
                root.ovrbtn(r.Overbeaten),
                root.btn(r.Beaten),
                r.HorseAge,
                r.HorseWeight,
                r.Headgear,
                root.time(r.FinishTime),
                r.DecimalOdds,
                r.JockeyID,
                r.TrainerID,
                root.prize(r.PrizeMoney),
                root.rating(r.OfficialRating),
                root.rating(r.RPRRating),
                root.rating(r.TSRating),
                r.OwnerID,
                r.Comment)

        if err != nil {
                _ = tx.Rollback()
                return err
        }

        if err := tx.Commit(); err != nil {
                return err
        }

        return nil
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
                case "":     return ""
                case "PU":   return "-1" // (Pulled up i.e. injury/issue)
                case "UR":   return "-2" // (Unseated Rider)
                case "DSQ":  return "-3" // (Disqualified)
                case "SU":   return "-4" // ?
                case "F":    return "-5" // (Fell)
                case "RR":   return "-6" // (Refused to Race)
                case "BD":   return "-7" // (Brought down)
                case "LFT":  return "-8" // ?
                case "RO":   return "-9" // (Refused to Race?)
        }
}

func (p *root) draw(s string) string {
        if s == "" {
                return "0" // blank?
        }
        return s
}

func (p *root) ovrbtn(s string) string {
        if s == "-" {
                return "0"
        }
        return s
}

func (p *root) btn(s string) string {
        if s == "-" {
                return "0"
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
                return "0"
        }
        return s
}

func (p *root) rating(s string) string {
        if s == "" || s == "–" {
                return "0"
        }
        return s
}

