package main 

import (
  "os"
  "os/exec"
  "encoding/json"
  "strconv"
  "fmt"

  "github.com/urfave/cli/v2"
)

var filepath string

func main() {
        app := &cli.App{
                Name: "Racingpost",
                Usage: "Download data from https://www.racingpost.com/",
                Flags: 
                        []cli.Flag{
                                &cli.StringFlag{
                                        Name: "filepath",
                                        Aliases: []string{"f"},
                                        Usage: "path to json containing params of what to download",
                                        Destination: &filepath,
                                },
                        },
                Action: func(*cli.Context) error {                        
                        return run()                
                },
        }

        if err := app.Run(os.Args); err != nil {
                panic(err)
        }
}

func run() error {
        fmt.Println("Scraping...")

        js, err := os.ReadFile(filepath)
        if err != nil {
                return err
        }

        var courses map[int]interface{}
        err = json.Unmarshal(js, &courses)
        if err != nil {
                return err
        }

        yearRange := []int{2008,2009,2010,2011,2012,2013,2014,2015,2016,2017,2018,2019,2020,2021,2022,2023}
        for id, c := range courses {
                fmt.Println(c)
                // convenient place to move scraped data
                fp := fmt.Sprintf("tools/rpscrape/data/courses/%s/flat", c)
                err = os.MkdirAll(fp, 0755)
                if err != nil {
                        return err
                }

                for _, yr := range yearRange {
                        fmt.Println(yr)
                        cmd := exec.Command("./rpscrape.py", 
                                "-c", strconv.Itoa(id), 
                                "-y", strconv.Itoa(yr), 
                                "-t", "flat")

                        cmd.Dir = "/home/ryan/dev/horse_racing/tools/rpscrape/scripts/" 

                        if err := cmd.Run(); err != nil {
                                return err
                        }

                        err = os.Rename(
                                fmt.Sprintf("tools/rpscrape/data/all/flat/%d.csv", yr), 
                                fmt.Sprintf("%s/%d.csv", fp, yr))

                        if err != nil {
                                return err
                        }
                }
        }

        return nil
}

