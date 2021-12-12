package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/dcoxall/aoc2021/dayone"
	"github.com/dcoxall/aoc2021/daythree"
	"github.com/dcoxall/aoc2021/daytwo"
)

func fail(str string) {
	fmt.Fprintf(os.Stderr, "Failed: %s\n", str)
	os.Exit(-1)
}

func skip() { fail("Day has been skipped") }

func withInput(filename string, partOne func(io.Reader) (int64, error), partTwo func(io.Reader) (int64, error)) {
	f, err := os.Open(filename)
	if err != nil {
		fail(err.Error())
	}

	var result int64
	start := time.Now()
	result, err = partOne(f)
	elapsed := time.Since(start)
	if err != nil {
		fail(err.Error())
	}

	fmt.Fprintf(os.Stdout, "Part One: %d (%s)\n", result, elapsed)

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		fail(err.Error())
	}

	start = time.Now()
	result, err = partTwo(f)
	elapsed = time.Since(start)
	if err != nil {
		fail(err.Error())
	}

	fmt.Fprintf(os.Stdout, "Part Two: %d (%s)\n", result, elapsed)
}

func day01() {
	withInput("input/01.txt", dayone.DepthChanges, dayone.SlidingDepthChanges)
}

func day02() {
	withInput("input/02.txt", daytwo.PartOne, daytwo.PartTwo)
}

func day03() {
	withInput("input/03.txt", daythree.PartOne, daythree.PartTwo)
}

var (
	targetDay   uint
	solutionMap = map[uint]func(){1: day01, 2: day02, 3: day03}
)

func init() {
	var lastDay uint
	for day := range solutionMap {
		if day > lastDay {
			lastDay = day
		}
	}

	flag.UintVar(&targetDay, "d", lastDay, "The day in which to run the test")
	flag.Parse()
}

func main() {
	var lastDay uint
	for day := range solutionMap {
		if day > lastDay {
			lastDay = day
		}
	}

	if solution, ok := solutionMap[targetDay]; ok {
		solution()
	} else {
		fail(fmt.Sprintf("Day %d has no available solution", targetDay))
	}
}
