package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

// Write persists mined and modelled data
type Write struct {
	db      *sql.DB
	convert *convert
}

// NewWrite creates a Write
func NewWrite(db *sql.DB) *Write {
	return &Write{
		db:      db,
		convert: &convert{},
	}

}

// Model persists a complete parse of a racing post record
func (wr *Write) Model(r *RacingPostRecord) error {
	// Write to all
	err_own := wr.Owner(r)
	err_jky := wr.Jockey(r)
	err_trn := wr.Trainer(r)
	err_hrs := wr.Horse(r)
	err_rce := wr.Race(r)
	err_run := wr.Runner(r)

	// Join all errors
	errs := errors.Join(err_own, err_jky, err_trn, err_hrs, err_rce, err_run)

	return errs
}

// RacingpostCSV persists mined Racingpost data
func (w *Write) RacingPost(cid int, year int, csv *[]byte) error {
	// Begin Transaction
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}

	// Create statement
	st := fmt.Sprintf("INSERT INTO racingpost VALUES($1, $2, $3, $4) ON CONFLICT DO NOTHING;")

	// Execute statement
	_, err = tx.Exec(st, cid, year, csv, false)

	// Rollback on error
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// Commit
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// Processed updates a racingpost record as now processed
func (w *Write) Processed(cid int, year int) error {
	// Begin Transaction
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}

	// Create statement
	st := fmt.Sprintf("UPDATE racingpost SET processed = true WHERE course_id = $1 AND year = $2;")

	// Execute statement
	_, err = tx.Exec(st, cid, year)

	// Rollback on error
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// Commit
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// Owner persists a race horse owner
func (w *Write) Owner(r *RacingPostRecord) error {
	// Begin Transaction
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}

	// Create statement
	st := fmt.Sprintf("INSERT INTO owner VALUES($1, $2, $3) ON CONFLICT DO NOTHING;")

	// Execute statement
	_, err = tx.Exec(st, r.OwnerID, r.OwnerName, r.SilkURL)

	// Rollback statement
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// Commit
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// Jockey persists a jockey
func (w *Write) Jockey(r *RacingPostRecord) error {
	// Begin transaction
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}

	// Create statemnt
	st := fmt.Sprintf("INSERT INTO jockey VALUES($1, $2) ON CONFLICT DO NOTHING;")

	// Execute statement
	_, err = tx.Exec(st, r.JockeyID, r.JockeyName)

	// Rollback on error
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// Commit
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// Trainer persists a race horse trainer
func (w *Write) Trainer(r *RacingPostRecord) error {
	// Begin transaction
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}

	// Create statement
	st := fmt.Sprintf("INSERT INTO trainer VALUES($1, $2) ON CONFLICT DO NOTHING;")

	// Execute statement
	_, err = tx.Exec(st, r.TrainerID, r.TrainerName)

	// Rollback on error
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// Commit
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// Horse persists a race horse
func (w *Write) Horse(r *RacingPostRecord) error {
	// Begin transaction
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}

	// Add Dam, assume gender not (F) filly
	_, err_d := tx.Exec("INSERT INTO horse VALUES($1, $2, $3) ON CONFLICT DO NOTHING;",
		r.DamID, r.DamName, "M")

	// Add Damsire, assume gender not (C) colt
	_, err_ds := tx.Exec("INSERT INTO horse VALUES($1, $2, $3) ON CONFLICT DO NOTHING;",
		r.DamsireID, r.DamsireName, "H")

	// Add Sire, assume gender not (C) colt
	_, err_s := tx.Exec("INSERT INTO horse VALUES($1, $2, $3) ON CONFLICT DO NOTHING;",
		r.SireID, r.SireName, "H")

	// Add Horse
	_, err_h := tx.Exec("INSERT INTO horse VALUES($1, $2, $3, $4, $5, $6) ON CONFLICT (horse_id) DO UPDATE SET name = $2, sex = $3, dam_id = $4, damsire_id = $5, sire_id = $6;",
		r.HorseID, r.HorseName, r.HorseSex, r.DamID, r.DamsireID, r.SireID)

	// Merge errors
	errs := errors.Join(err_d, err_ds, err_s, err_h)

	// Rollback on error
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	// Commit
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// Race persists a completed race
func (w *Write) Race(r *RacingPostRecord) error {
	// Begin transaction
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}

	// Create statement
	st := fmt.Sprintf("INSERT INTO race VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) ON CONFLICT DO NOTHING;")

	// Execute statement
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

	// Rollback on error
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// Commit
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// Race persists a runner of a completed race
func (w *Write) Runner(r *RacingPostRecord) error {
	// Begin statement
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}

	// Create statement
	st := fmt.Sprintf("INSERT INTO runner VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20) ON CONFLICT DO NOTHING;")

	// Execute statement
	_, err = tx.Exec(st,
		r.HorseID,
		r.RaceID,
		w.convert.num(r),
		w.convert.position(r),
		w.convert.draw(r),
		w.convert.ovrbtn(r),
		w.convert.btn(r),
		r.HorseAge,
		r.HorseWeight,
		r.Headgear,
		w.convert.time(r),
		r.DecimalOdds,
		r.JockeyID,
		r.TrainerID,
		w.convert.prize(r),
		w.convert.rating(r.OfficialRating),
		w.convert.rating(r.RPRRating),
		w.convert.rating(r.TSRating),
		r.OwnerID,
		r.Comment)

	// Rollback on error
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// Commit
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// RacingpostRecord to model conversions
type convert struct{}

func (c *convert) num(r *RacingPostRecord) string {
	switch r.RacecardNumber {
	default:
		return r.RacecardNumber
	case "":
		return "0" // blank?
	}
}

// Abbreviations help - https://help.racingpost.com/hc/en-us/articles/115001699689-Abbreviations-on-the-racecard
func (c *convert) position(r *RacingPostRecord) string {
	switch r.FinishedPosition {
	default:
		return r.FinishedPosition
	case "":
		return ""
	case "PU":
		return "-1" // (Pulled up i.e. injury/issue)
	case "UR":
		return "-2" // (Unseated Rider)
	case "DSQ":
		return "-3" // (Disqualified)
	case "SU":
		return "-4" // ?
	case "F":
		return "-5" // (Fell)
	case "RR":
		return "-6" // (Refused to Race)
	case "BD":
		return "-7" // (Brought down)
	case "LFT":
		return "-8" // ?
	case "RO":
		return "-9" // (Refused to Race?)
	}
}

func (c *convert) draw(r *RacingPostRecord) string {
	switch r.Draw {
	default:
		return r.Draw
	case "":
		return "0" // blank?
	}
}

func (c *convert) ovrbtn(r *RacingPostRecord) string {
	switch r.Overbeaten {
	default:
		return r.Overbeaten
	case "-":
		return "0"
	}
}

func (c *convert) btn(r *RacingPostRecord) string {
	switch r.Beaten {
	default:
		return r.Beaten
	case "-":
		return "0"
	}
}

func (c *convert) time(r *RacingPostRecord) *string {
	switch r.FinishTime {
	case "-":
		return nil
	default:
		return &r.FinishTime
	}
}

func (c *convert) prize(r *RacingPostRecord) string {
	switch r.PrizeMoney {
	case "", "–":
		return "0"
	default:
		return r.PrizeMoney
	}
}

func (c *convert) rating(s string) string {
	switch s {
	case "", "–":
		return "0"
	default:
		return s
	}
}
