package dayone

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func DepthChanges(rdr io.Reader) (int64, error) {
	// We want to store the number of depth changes and the
	// previously seen depth (for comparison). I am also
	// initializing the current depth value outside the loop.
	var (
		depthChanges  int64 = -1 // The first depth is not a change
		currentDepth  int64
		previousDepth int64

		err error
	)

	// Iterate through each line of the input
	scanner := bufio.NewScanner(rdr)
	for scanner.Scan() {
		currentDepth, err = strconv.ParseInt(scanner.Text(), 10, 64)

		if err != nil {
			return -1, err
		}

		if currentDepth > previousDepth {
			depthChanges += 1
		}

		previousDepth = currentDepth
	}

	if err = scanner.Err(); err != nil {
		// Wrap the error
		err = fmt.Errorf("Failed to read input: %w", err)
	}

	return depthChanges, err
}

func asyncParseInput(inputs io.Reader, nums chan<- int64, errs chan<- error, finish chan<- bool) {
	var (
		// Reserve memory for the value we need to parse
		current int64
		// We could encounter an error so lets reserve memory for that too
		parseErr error
	)

	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		current, parseErr = strconv.ParseInt(scanner.Text(), 10, 64)

		// If we encounter an error send it via the errs channel
		if parseErr != nil {
			errs <- parseErr
			break
		} else {
			// Otherwise we can send the parsed number to the nums channel
			nums <- current
		}
	}

	// If we broke out the loop or even finished parsing we still need to check
	// if the scanner errored. If so this is also sent to the errs channels
	if parseErr = scanner.Err(); parseErr != nil {
		errs <- parseErr
	}

	// Now we are done so send finish and close channels
	finish <- true
	close(errs)
	close(nums)
	close(finish)
}

func SlidingDepthChanges(rdr io.Reader) (int64, error) {
	// This time we will need to parse ahead 3 lines at a time
	// and iterate over line indices, again tracking the depth
	// changes but the current is determined by suming
	// 3 measurements inclusive of the current position
	var (
		depthChanges  int64   = -1 // The first depth is not a change
		currentValues []int64 = make([]int64, 0, 3)
		currentSum    int64
		previousSum   int64
	)

	// To perform just-in-time parsing I will move it to a buffered
	// channel with a limit of 3
	parsed := make(chan int64, 3)
	// I also need a way to communicate a failure to parse
	errs := make(chan error, 1)
	// Finally yet another channel to indicate completeness
	finish := make(chan bool, 1)

	go asyncParseInput(rdr, parsed, errs, finish)

	for {
		exit := false

		select {
		case err := <-errs:
			return depthChanges, fmt.Errorf("Failed to parse input: %w", err)
		case n := <-parsed:
			if len(currentValues) == 3 {
				// We need to remove the first element and append the new element
				currentValues = append(currentValues[1:], n)
			} else {
				currentValues = append(currentValues, n)
			}
		case <-finish:
			// This is used so we can exit the surrounding for loop
			// without using named breaks
			exit = true
		}

		if len(currentValues) == 3 {
			// we are full so we want to work out the sum for this window
			currentSum = 0
			for _, v := range currentValues {
				currentSum += v
			}
			// lets count it if this window was greater than the last
			if currentSum > previousSum {
				depthChanges += 1
			}
			// assign the new sum as the previous
			previousSum = currentSum
		}

		if exit {
			break
		}
	}

	return depthChanges, nil
}
