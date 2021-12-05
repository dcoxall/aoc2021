package daytwo

import (
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
	testInput, err := os.Open("testinput.txt")
	if err != nil {
		t.Fatal(err)
	}

	var result int64
	result, err = PartOne(testInput)
	if err != nil {
		t.Fatal(err)
	} else if result != 150 {
		t.Fatalf("Expected 150 but got %d", result)
	}
}

func TestPartTwo(t *testing.T) {
	testInput, err := os.Open("testinput.txt")
	if err != nil {
		t.Fatal(err)
	}

	var result int64
	result, err = PartTwo(testInput)
	if err != nil {
		t.Fatal(err)
	} else if result != 900 {
		t.Fatalf("Expected 900 but got %d", result)
	}
}
