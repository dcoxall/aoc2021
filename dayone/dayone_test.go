package dayone

import (
	"os"
	"testing"
)

func TestDepthChanges(t *testing.T) {
	testInput, err := os.Open("testinput.txt")
	if err != nil {
		t.Fatal(err)
	}

	var result int64
	result, err = DepthChanges(testInput)
	if err != nil {
		t.Fatal(err)
	} else if result != 7 {
		t.Fatalf("Expected 7 but got %d", result)
	}
}

func TestSlidingDepthChanges(t *testing.T) {
	testInput, err := os.Open("testinput.txt")
	if err != nil {
		t.Fatal(err)
	}

	var result int64
	result, err = SlidingDepthChanges(testInput)
	if err != nil {
		t.Fatal(err)
	} else if result != 5 {
		t.Fatalf("Expected 5 but got %d", result)
	}
}
