package main

import (
  "database/sql"
  "os"
  "io"

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
                                Name: "model",
                                Usage: "Transform racingpost data to model",
                                Flags: []cli.Flag{
                                        &cli.IntFlag{
                                                Name: "course-id",
                                                Aliases: []string{"c", "cid"},
                                                Usage: "course id of csv to transform to model",
                                                Destination: &cid,
                                        },
                                        &cli.IntFlag{
                                                Name: "year",
                                                Aliases: []string{"y"},
                                                Usage: "year of csv to transform to model",
                                                Destination: &year,
                                        },
                                },
                                Action: func (cCtx *cli.Context) error {
                                        // Check for empty required params

                                        if cid == 0 {
                                                panic("course-id empty")
                                        }
                                        if year == 0 {
                                                panic("year empty")
                                        }

                                        return model(cid, year)
                                },
                        },
                        {
                                Name: "csv",
                                Usage: "Add a csv",
                                Flags: []cli.Flag{
                                        &cli.StringFlag{
                                                Name: "file",
                                                Aliases: []string{"f"},
                                                Usage: "path to csv file",
                                                Destination: &filepath,
                                        },
                                        &cli.IntFlag{
                                                Name: "course-id",
                                                Aliases: []string{"c","cid"},
                                                Usage: "the course id of data",
                                                Destination: &cid,
                                        },
                                        &cli.IntFlag{
                                                Name: "year",
                                                Aliases: []string{"y"},
                                                Usage: "the year the data was recorded",
                                                Destination: &year,
                                        },
                                },
                                Action: func (cCtx *cli.Context) error {
                                        // Check for empty required params

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
                panic(err)
        }
}

func racingpost(cid int, year int, filepath string) error {
        // Open DB connection
        db, err := sql.Open("postgres", "postgresql://localhost/horse_racing?sslmode=disable")
        if err != nil {
                panic(err)
        }
        defer db.Close()

        // Open CSV file
        csv, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0755)
        if err != nil {
                return err
        }
        defer csv.Close()

        // Read all CSV contents
        cnt, err := io.ReadAll(csv)
        if err != nil {
                return err
        }

        // Create DB writer
        write := NewWrite(db)

        // Create RacingPost record
        if err := write.RacingPost(cid, year, &cnt); err != nil {
                return err
        }

        return nil
}

func model(cid int, year int) error {
        // Open DB connection
        db, err := sql.Open("postgres", "postgresql://localhost/horse_racing?sslmode=disable")
        if err != nil {
                panic(err)
        }
        defer db.Close()

        // Create DB reader
        read := Read{ db }

        // Query RacingPost records
        records, err := read.RacingPost(cid, year)
        if err != nil {
                return err
        }

        // Create DB writer
        write := NewWrite(db)

        for _, r := range records {

                // Create/Transform to model
                if err := write.Model(r); err != nil {
                        return err
                }

                // Update Racingpost record as processed
                // TODO move this to transaction within write.Model
                if err := write.Processed(cid, year); err != nil {
                        return err
                }
        }

        return nil
}

