package main

import (
  // "fmt"
  "database/sql"
  "os"

  // "github.com/rfrenchy/punts/cmd/repository/internal/write"

  "github.com/urfave/cli/v2"
  "github.com/rs/zerolog"
  "github.com/rs/zerolog/log"
  _ "github.com/lib/pq"
)

func main() {
        zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
        // log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

        var filepath string

        app := &cli.App{
                Commands: []*cli.Command{
                        {
                                Name: "owner",
                                Aliases: []string{"o"},
                                Usage: "Add an owner",
                                Action: func (cCtx *cli.Context) error {
                                        return addOwner(filepath)
                                },
                        },
                },
        }

        if err := app.Run(os.Args); err != nil {
                log.Panic().Err(err).Msg("panicing")
                panic(err)
        }
}

func run(filepath string) error {
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

        return nil

        //return write.RacingPost(csv)
}

func addOwner(filepath string) error {
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

        write := &Write{ db: db, }
        record := &RacingPostRecord{}

        return write.Owner(record)
}

