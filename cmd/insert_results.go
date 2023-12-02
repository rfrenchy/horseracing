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
  "github.com/rs/zerolog"
  "github.com/rs/zerolog/log"
  "github.com/gocarina/gocsv"
  _ "github.com/lib/pq"
)

// Racingpost CSV Record
type Record struct {
        // Race information
        RaceDate string                 `csv:"date"`
        RaceRegion string               `csv:"region"`
        CourseID string                 `csv:"course_id"`
        Course string                   `csv:"course"`
        RaceID string                   `csv:"race_id"`
        RaceOfftime string              `csv:"off"`
        RaceName string                 `csv:"race_name"`
        Racetype string                 `csv:"type"`
        Raceclass string                `csv:"class"`
        RacePattern string              `csv:"pattern"` // (Group 1 etc)
        RatingbandRestrictions string   `csv:"rating_band"`
        AgebandRestriction string       `csv:"age_band"`
        SexRestriction string           `csv:"sex_rest"`
        Distance string                 `csv:"dist_m"` // in metres
        Going string                    `csv:"going"` // going description
        Surface string                  `csv:"surface"` // surface Turf/Dirt/AW

        // Specific Runner information
        RacecardNumber string   `csv:"num"`
        FinishedPosition string `csv:"position"` // runner finished position
        Draw string             `csv:"draw"` // stall number
        Overbeaten string       `csv:"ovr_btn"` // total number of lengths beaten
        Beaten string           `csv:"btn"` // lengths behind nearest horse in front
        HorseID string          `csv:"horse_id"`
        HorseName string        `csv:"horse"`
        HorseAge string         `csv:"age"`
        HorseSex string         `csv:"sex"`
        HorseWeight string      `csv:"lbs"` // weight in pounds
        Headgear string         `csv:"hg"`
        FinishTime string       `csv:"time"` // time taken in minutes/seconds
        DecimalOdds string      `csv:"dec"`
        JockeyId string         `csv:"jockey_id"`
        Jockey string           `csv:"jockey"`
        TrainerId string        `csv:"trainer_id"`
        Trainer string          `csv:"trainer"`
        PrizeMoney string       `csv:"prize"`
        OfficialRating string   `csv:"or"`
        RPRRating string        `csv:"rpr"`
        TSRating string         `csv:"ts"`
        SireID string           `csv:"sire_id"`
        SireName string         `csv:"sire"`
        DamID string            `csv:"dam_id"`
        DamName string          `csb:"dam"`
        DamsireID string        `csv:"damsire_id"`
        DamsireName string      `csv:"damsire"`
        OwnerID string          `csv:"owner_id"`
        OwnerName string        `csv:"owner"`
        SilkURL string          `csv:"silk_url"` // URL of silk colours
        Comment string          `csv:"comment"` // Form in running comments
}

type Race struct {
        Id int
        Name string
        Date string // date of race
        Region string // region of race
        Course Course
        Offtime string // race off time
        Racetype string // type of racing (flat/hurdle/chase etc)
        Raceclass string // race class
        Pattern string // race pattern
        Ratingband string // rating restrictions
        Agebandrestriction string // age restrictions
        Sexrestriction string // sex restrictions
        Distance string // in metres
        Going string // going description
        Surface string // surface turf/dirt/aw
        Ran int // number of runners in race
}

type Course struct {
        Id int
        Name string
        Region int
}

type Region struct {
        Id int
        Name string
}

type Runner struct {
        Horse int
        Race int
        Racecardnumber int
        Position int // finished position
        Draw int // stall number
        Overbeaten int // total number of lengths beaten
        Beaten int // lengths behind nearest horse in front
        Age string // age of horse at race time
        Weight string // weight in pounds
        Headgear string
        Time string // time in minutes/seconds
        Odds string // decimal odds
        Jockeyid int
        Trainerid int
        Prizemoney int
        Officialrating int
        RPRating int
        TSrating int
        Ownerid int
        Comment string
}

type Horse struct {
        ID int
        Name string
        SireID int // Father
        DamID int // Mother
        DamsireID int // Father of the Dam
        Sex string
}

type Jockey struct {
        Id int
        Name string
}

type Trainer struct {
        Id int
        Name string
}

type Owner struct {
        Id int
        Name string
        Silkurl string
}

func init() {
}

func main() {
        zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
        // log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

        var commit bool // TODO change to string flag  (raw | lingfield | result)
        var filepath string

        app := &cli.App{
                Name: "Inserter-Racingpost",
                Usage: "Extracts and processes horse-racing CSV, inserts into PostresDB",
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
                log.Panic().Err(err).Msg("panicing")
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
        defer db.Close()

        csv, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0755)
        if err != nil {
                return err
        }
        defer csv.Close()

	records := []*Record{}
	if err := gocsv.UnmarshalFile(csv, &records); err != nil {
		return err
	}

        // CREATE INSERT STATEMENT STRING
        //var sb strings.Builder
        //sb.WriteString("$" + strconv.Itoa(i + 1) + ",")
        //v := strings.TrimRight(sb.String(), ",") // remove trailing comma

        // To Domain objects
        // start at 1, skip header line
        for i := 1; i < len(records); i++ {
                r := records[i]

                h := &Horse{}

                if id, err := strconv.Atoi(r.HorseID); err != nil {
                        h.ID = nil
                } else {
                        h.ID = id
                }

//                h := &Horse{
//                        ID: strconv.Atoi(r.HorseID),
//                        Name: r.HorseName,
//                        SireID: strconv.Atoi(r.SireID),
//                        DamID:  strconv.Atoi(r.DamID),
//                        DamsireID: strconv.Atoi(r.DamsireID),
//                        Sex: r.HorseSex,
//                }
//

        }


        return nil

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


        pr := &root{}

                for i, x := range record {
                        // $1, $2, $3... etc
                        switch i {
                                default: p[i] = x
                                case 16: p[i] = pr.num(x)
                                case 17: p[i] = pr.position(x)
                                case 18: p[i] = pr.draw(x)
                                case 19: p[i] = pr.ovrbtn(x)
                                case 20: p[i] = pr.btn(x)
                                case 26: p[i] = pr.time(x)
                                case 27: p[i] = pr.seconds(x)
                                case 31: p[i] = pr.prize(x)
                                case 33: p[i] = pr.rpr(x)
                        }
                }

}

type root struct{}

// type root struct {}
// type records struct { horse[], runner[], race[], course[]}
// func Create(CSVRecord) records
// type insert struct{}
// func *insert Horses(horse[])
// func *insert Runners(runner[])
// func *inesrt Races(race[])
// etc...


// TODO, change to string and/or nil?
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

// TODO, change to string and/or nil?
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

