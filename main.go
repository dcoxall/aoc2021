package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/dcoxall/aoc2021/dayone"
)

func fail(str string) {
	fmt.Fprintf(os.Stderr, "Failed: %s\n", str)
	os.Exit(-1)
}

func skip() { fail("Day has been skipped") }

func day01() {
	// Read the input
	f, err := os.Open("input/01.txt")
	if err != nil {
		fail(err.Error())
	}

	var result int64
	result, err = dayone.DepthChanges(f)
	if err != nil {
		fail(err.Error())
	}

	fmt.Fprintf(os.Stdout, "Depth changes: %d\n", result)

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		fail(err.Error())
	}
	result, err = dayone.SlidingDepthChanges(f)
	if err != nil {
		fail(err.Error())
	}

	fmt.Fprintf(os.Stdout, "Sliding depth changes: %d\n", result)
}

var solutionMap = map[uint]func(){1: day01}

func main() {
	var lastDay uint
	for day := range solutionMap {
		if day > lastDay {
			lastDay = day
		}
	}

	targetDay := flag.Uint("d", lastDay, "The day in which to run the test")

	if solution, ok := solutionMap[*targetDay]; ok {
		solution()
	} else {
		fail(fmt.Sprintf("Day %d is has no available solution", targetDay))
	}
}
