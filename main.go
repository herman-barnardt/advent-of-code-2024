package main

import (
	_ "advent-of-code-2024/day1"
	_ "advent-of-code-2024/day2"
	_ "advent-of-code-2024/day3"
	_ "advent-of-code-2024/day4"
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/herman-barnardt/aoc"
)

func main() {
	flag.Parse()

	command := flag.Arg(0)
	year := 2024
	_, _, day := time.Now().Date()
	var err error
	dayString := flag.Arg(1)
	if len(dayString) > 0 && dayString != "0" {
		day, err = strconv.Atoi(dayString)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
	}
	part := 0
	partString := flag.Arg(2)
	if len(partString) > 0 {
		part, err = strconv.Atoi(partString)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
	}

	err = aoc.Run(command, year, day, part)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
