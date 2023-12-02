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
  _ "github.com/lib/pq"
)

func init() {
}

func main() {
        zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
        // log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

        var commit bool // TODO change to string flag  (raw | lingfield | result)
        var filepath string

        app := &cli.App{
                Name: "Punting",
                Usage: "",
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

        return write.RacingPostRecord(f)
}

