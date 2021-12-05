package daytwo

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Vector3D struct {
	X int64
	Y int64
	Z int64
}

func (a *Vector3D) Add(b *Vector3D) Vector3D {
	return Vector3D{X: a.X + b.X, Y: a.Y + b.Y, Z: a.Z + b.Z}
}

func parseInput(rdr io.Reader) ([]Vector3D, error) {
	input := make([]Vector3D, 0)
	scanner := bufio.NewScanner(rdr)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")  // Split the line into a direction and value
		n, err := strconv.ParseInt(parts[1], 10, 64) // Parse the value

		if err != nil {
			return input, err
		}

		switch parts[0] {
		case "up":
			input = append(input, Vector3D{X: 0, Y: n * -1, Z: 0})
		case "down":
			input = append(input, Vector3D{X: 0, Y: n * 1, Z: 0})
		case "forward":
			input = append(input, Vector3D{X: n, Y: 0, Z: 0})
		}
	}

	err := scanner.Err()
	return input, err
}

func PartOne(rdr io.Reader) (int64, error) {
	vectors, err := parseInput(rdr)
	if err != nil {
		return -1, err
	}

	position := Vector3D{X: 0, Y: 0, Z: 0}
	for _, vector := range vectors {
		position = position.Add(&vector)
	}

	return position.X * position.Y, nil
}

func PartTwo(rdr io.Reader) (int64, error) {
	vectors, err := parseInput(rdr)
	if err != nil {
		return -1, err
	}

	// We need to re-evaluate the vectors now to apply
	// the corrected calculation
	// X = horizontal, Y = aim, Z = depth
	position := Vector3D{X: 0, Y: 0, Z: 0}
	for _, vector := range vectors {
		position = position.Add(&vector)
		if vector.X != 0 {
			// now additionally apply the new logic
			position.Z += position.Y * vector.X
		}
	}

	return position.X * position.Z, nil
}
