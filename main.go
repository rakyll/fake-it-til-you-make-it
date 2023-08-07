package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"time"
)

const WhatTheCommit = "https://whatthecommit.com/index.txt"

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	max         int
	skipWeekend bool
	randomMsg   bool
	years       int
)

// getRandomMsg returns random commit message
func getRandomMsg() string {
	resp, err := http.Get(WhatTheCommit)
	if err != nil {
		return "something wrong"
	}

	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

func main() {
	flag.IntVar(&max, "max", 5, "Max number of commits per day")
	flag.BoolVar(&skipWeekend, "skip-weekend", true, "Set false if you don't want to skip weekend")
	flag.BoolVar(&randomMsg, "random-msg", false, "Set random commit message, it may slow down the program")
	flag.IntVar(&years, "years", 2, "Number of years to commit")
	flag.Parse()

	end := time.Now()
	start := end.AddDate(-1*years, 0, 0)

	for d := start; d.Before(end); d = d.AddDate(0, 0, 1) {
		if skipWeekend {
			if d.Weekday() == time.Saturday || d.Weekday() == time.Sunday {
				continue
			}
		}

		commits := rand.Intn(max)
		log.Printf("Committing %d times on %s", commits, d)

		for i := 0; i < commits; i++ {
			msg := fmt.Sprintf("Commit from %s: ", d.Format("2006-01-02"))
			if randomMsg {
				msg = getRandomMsg()
			}

			cmd := exec.Command(
				"git",
				"commit",
				"--allow-empty",
				"--date", d.Format("Mon Jan 02 15:04:05 -0700 2006"),
				"-m", msg)
			if err := cmd.Run(); err != nil {
				log.Printf("Error when committing: %v", err)
			}
		}
	}
}
