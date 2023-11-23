package main 

import (
  "os"
  "os/exec"
  "encoding/json"
  "strconv"
  "fmt"
)

func main() {
        fmt.Println("Scraping...")

        js, err := os.ReadFile("./scripts/flat.json")
        if err != nil {
                panic(err)
        }

        var courses map[int]interface{}
        err = json.Unmarshal(js, &courses)
        if err != nil {
                panic(err)
        }

        yearRange := []int{2008,2009,2010,2011,2012,2013,2014,2015,2016,2017,2018,2019,2020,2021,2022,2023}
        for id, c := range courses {
                fmt.Println(c)
                // convenient place to move scraped data
                fp := fmt.Sprintf("tools/rpscrape/data/courses/%s/flat", c)
                err = os.MkdirAll(fp, 0755)
                if err != nil {
                        panic(err)
                }

                for _, yr := range yearRange {
                        fmt.Println(yr)
                        cmd := exec.Command("./rpscrape.py", 
                                "-c", strconv.Itoa(id), 
                                "-y", strconv.Itoa(yr), 
                                "-t", "flat")

                        cmd.Dir = "/home/ryan/dev/horse_racing/tools/rpscrape/scripts/" 

                        if err := cmd.Run(); err != nil {
                                panic(err)
                        }

                        err = os.Rename(
                                fmt.Sprintf("tools/rpscrape/data/all/flat/%d.csv", yr), 
                                fmt.Sprintf("%s/%d.csv", fp, yr))

                        if err != nil {
                                panic(err)
                        }
                }
        }
}

