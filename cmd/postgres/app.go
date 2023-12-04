package db

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

func main() {
        zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
        // log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

        var filepath string

        app := &cli.App{
                Name: "db",
                Usage: "",
                Action: func(*cli.Context) error {
                        return run(filepath)
                },
                Flags:
                        []cli.Flag{
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

func run(filepath string) error {
        // Open DB connection
        // TODO load from config
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

        return write.RacingPost(f)
}

