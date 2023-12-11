package main

import (
  // "fmt"
  "database/sql"
  "os"

  "github.com/gocarina/gocsv"
  "github.com/urfave/cli/v2"
  "github.com/rs/zerolog"
  "github.com/rs/zerolog/log"
  _ "github.com/lib/pq"
)

var owner bool
var jockey bool
var trainer bool
var horse bool
var race bool
var runner bool

func main() {
        zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
        // log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})


        var filepath string

        app := &cli.App{
                Commands: []*cli.Command{
                        {
                                Name: "add",
                                Usage: "Add a record",
                                Flags: []cli.Flag{
                                        &cli.StringFlag{
                                                Name: "file",
                                                Aliases: []string{"f"},
                                                Usage: "path to racingpost csv file",
                                                Destination: &filepath,
                                        },
                                        &cli.BoolFlag{
                                                Name: "owner",
                                                Aliases: []string{"own"},
                                                Usage: "Add owners",
                                                Destination: &owner,
                                        },
                                        &cli.BoolFlag{
                                                Name: "jockey",
                                                Aliases: []string{"jky"},
                                                Usage: "Add jockeys",
                                                Destination: &jockey,
                                        },
                                        &cli.BoolFlag{
                                                Name: "trainer",
                                                Aliases: []string{"trn"},
                                                Usage: "Add trainers",
                                                Destination: &trainer,
                                        },
                                        &cli.BoolFlag{
                                                Name: "horse",
                                                Aliases: []string{"hrs"},
                                                Usage: "Add horse",
                                                Destination: &horse,
                                        },
                                        &cli.BoolFlag{
                                                Name: "race",
                                                Aliases: []string{"rce"},
                                                Usage: "Add a race",
                                                Destination: &race,
                                        },
                                        &cli.BoolFlag{
                                                Name: "runner",
                                                Aliases: []string{"rnr"},
                                                Usage: "Add a runner",
                                                Destination: &runner,
                                        },
                                },
                                Action: func (cCtx *cli.Context) error {
                                        if filepath == "" {
                                                panic("path to file required")
                                        }

                                        return add(filepath)
                                },
                        },
                },
        }

        if err := app.Run(os.Args); err != nil {
                log.Panic().Err(err).Msg("panicing")
                panic(err)
        }
}

func add(filepath string) error {
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
                if owner {
                        if err := write.Owner(r); err != nil {
                                return err
                        }
                }

                if jockey {
                        if err := write.Jockey(r); err != nil {
                                return err
                        }
                }

                if trainer {
                        if err := write.Trainer(r); err != nil {
                                return err
                        }
                }

                if horse {
                        if err := write.Horse(r); err != nil {
                                return err
                        }
                }

                if race {
                        if err := write.Race(r); err != nil {
                                return err
                        }
                }

                if runner {
                        if err := write.Runner(r); err != nil {
                                return err
                        }
                }
        }

        return nil
}

