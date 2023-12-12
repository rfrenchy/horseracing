package main

import (
  "database/sql"
  "os"
  "io"

  "github.com/gocarina/gocsv"
  "github.com/urfave/cli/v2"
  "github.com/rs/zerolog"
  "github.com/rs/zerolog/log"
  _ "github.com/lib/pq"
)

func main() {
        zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
        // log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

        var filepath string
        var cid int
        var year int

        app := &cli.App{
                Commands: []*cli.Command{
                        {
                                Name: "transform",
                                Usage: "Transform racingpost data to model",
                                Flags: []cli.Flag{
                                        &cli.StringFlag{
                                                Name: "--course-id",
                                                Aliases: []string{"c", "cid"},
                                                Usage: "course id of csv to transform to model",
                                                Destination: &cid,
                                        },
                                        &cli.StringFlag{
                                                Name: "--year",
                                                Aliases: []string{"y"},
                                                Usage: "year of csv to transform to model",
                                                Destination: &year,
                                        },
                                },
                                Action: func (cCtx *cli.Context) error {
                                        if cid == 0 {
                                                panic("course-id empty")
                                        }
                                        if year == 0 {
                                                panic("year empty")
                                        }

                                        return add(filepath)
                                },
                        },
                        {
                                Name: "csv",
                                Usage: "Add a csv",
                                Flags: []cli.Flag{
                                        &cli.StringFlag{
                                                Name: "--file",
                                                Aliases: []string{"f"},
                                                Usage: "path to csv file",
                                                Destination: &filepath,
                                        },
                                        &cli.IntFlag{
                                                Name: "--course-id",
                                                Aliases: []string{"c","cid"},
                                                Usage: "the course id of data",
                                                Destination: &cid,
                                        },
                                        &cli.IntFlag{
                                                Name: "--year",
                                                Aliases: []string{"y"},
                                                Usage: "the year belonging to the data",
                                                Destination: &year,
                                        },
                                },
                                Action: func (cCtx *cli.Context) error {
                                        if filepath == "" {
                                                panic("filepath empty")
                                        }
                                        if cid == 0 { // i.e empty
                                                panic("course id empty")
                                        }
                                        if year == 0 { // i.e. empty
                                                panic("year empty")
                                        }

                                        return racingpost(cid, year, filepath)
                                },
                        },
                },
        }

        if err := app.Run(os.Args); err != nil {
                log.Panic().Err(err).Msg("panicing")
                panic(err
        }
}

func racingpost(cid int, year int, filepath string) error {
        // Open DB connection
        db, err := sql.Open("postgres", "postgresql://localhost/horse_racing?sslmode=disable")
        if err != nil {
                panic(err)
        }
        defer db.Close()

        csv, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0755)
        if err != nil {
                return err
        }
        defer csv.Close()

        c, err := io.ReadAll(csv)
        if err != nil {
                return err
        }

        write := Write{ db: db }

        if err := write.Racingpost(cid, year, &c); err != nil {
                return err
        }

        return nil
}

func add(filepath string) error {
        // transform(cid int, year int) error
        // read racingpost cid, year
        // call write.Add with racingpost row


        // ********
        // * TODO *
        // ********

        // Open DB connection
        db, err := sql.Open("postgres", "postgresql://localhost/horse_racing?sslmode=disable")
        if err != nil {
                panic(err)
        }
        defer db.Close()

        csv, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0755)
        if err != nil {
                return err
        }
        defer csv.Close()

        write := Write{ db: db }

        records := []*RacingPostRecord{}
        if err := gocsv.UnmarshalFile(csv, &records); err != nil {
	        return err
	}

        for _, r := range records {
                if err := write.Add(r); err != nil {
                        return err
                }
        }

        return nil
}

