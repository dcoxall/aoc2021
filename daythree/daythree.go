package daythree

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func parseInput(rdr io.Reader) ([]uint64, int, error) {
	input := make([]uint64, 0)
	scanner := bufio.NewScanner(rdr)
	length := 0
	for scanner.Scan() {
		str := scanner.Text()
		if length == 0 {
			length = len(str)
		}
		str = fmt.Sprintf("%016s", str)
		n, err := strconv.ParseUint(str, 2, 16)
		if err != nil {
			return input, length, err
		}

		input = append(input, n)
	}
	return input, length, scanner.Err()
}

func PartOne(rdr io.Reader) (int64, error) {
	input, max, err := parseInput(rdr)
	if err != nil {
		return -1, err
	}

	// We calculate how many bits are needed to determine
	// the dominant bit (0 or 1)
	var midPoint uint64 = (uint64(len(input)) / 2) + 1

	// Next we need to iterate over each of the inputs and
	// begin counting the bits by postion... This is done
	// using a bit mask and bit shifting...
	// For positon 1 as an example:
	// 10110 & (1 << 1) = 10110 & 00010 = 00010
	// which is then shifted
	// 00010 >> 1 = 00001 = 1
	// is this all necessary? probably not but I don't get
	// to play with bit shifting in the day job
	counts := make(map[int]uint64)
	for _, n := range input {
		for i := 0; i < max; i++ {
			counts[i] += (n & (1 << i)) >> i
		}
	}

	// Now we calculate the gamma by comparing the count to the
	// mid point and if it is greater or equal then we can shift
	// OR the bit into the relevant position
	var (
		gamma uint64 = 0
		mask  uint64 = 0
	)
	for i, count := range counts {
		if midPoint <= count {
			gamma = gamma | 1<<i
		}
	}

	for i := 0; i < max; i++ {
		mask = mask | 1<<i
	}

	return int64(gamma * (gamma ^ mask)), nil
}

func mostAndLeastCommonByBit(coll []uint64, i int) ([]uint64, []uint64) {
	var (
		partitionA []uint64 = make([]uint64, 0) // everything with a 1 bit at the index
		partitionB []uint64 = make([]uint64, 0) // everything with a 0 bit at the index
		aLength    uint     = 0
		bLength    uint     = 0
	)
	for _, n := range coll {
		if (n & (1 << i)) == 0 {
			partitionB = append(partitionB, n)
			bLength += 1
		} else {
			partitionA = append(partitionA, n)
			aLength += 1
		}
	}

	if aLength >= bLength {
		return partitionA, partitionB
	} else {
		return partitionB, partitionA
	}
}

func PartTwo(rdr io.Reader) (int64, error) {
	input, max, err := parseInput(rdr)
	if err != nil {
		return -1, err
	}

	oxygenRatings, c02Ratings := mostAndLeastCommonByBit(input, max-1)

	for index := max - 2; index >= 0; index-- {
		if len(oxygenRatings) > 1 {
			oxygenRatings, _ = mostAndLeastCommonByBit(oxygenRatings, index)
		}
		if len(c02Ratings) > 1 {
			_, c02Ratings = mostAndLeastCommonByBit(c02Ratings, index)
		}
	}

	return int64(oxygenRatings[0] * c02Ratings[0]), nil
}
