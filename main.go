package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	max         int
	skipWeekend bool
	years       int
	months      int
)

func main() {
	flag.IntVar(&max, "max", 5, "Max number of commits per day")
	flag.BoolVar(&skipWeekend, "skip-weekend", true, "Set false if you don't want to skip weekend")
	flag.IntVar(&years, "years", 0, "Number of months to commit")
	flag.IntVar(&months, "months", 2, "Number of months to commit")
	flag.Parse()

	end := time.Now()
	start := end.AddDate(-1*years, -1*months, 0)

	for d := start; d.Before(end); d = d.AddDate(0, 0, 1) {
		if skipWeekend {
			if d.Weekday() == time.Saturday || d.Weekday() == time.Sunday {
				continue
			}
		}

		commits := rand.Intn(max)
		log.Printf("Committing %d times on %s", commits, d)

		for i := 0; i < commits; i++ {
			cmd := exec.Command(
				"git",
				"commit",
				"--allow-empty",
				"--date", d.Format("Mon Jan 02 15:04:05 -0700 2006"),
				"-m", fmt.Sprintf("Commit from %s", d.Format("2006-01-02")))
			if err := cmd.Run(); err != nil {
				log.Printf("Error when committing: %v", err)
			}
		}
	}
}
