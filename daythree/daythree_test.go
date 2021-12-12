package daythree

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
	} else if result != 198 {
		t.Fatalf("Expected 198 but got %d", result)
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
	} else if result != 230 {
		t.Fatalf("Expected 230 but got %d", result)
	}
}
