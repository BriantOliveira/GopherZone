package grains

import (
	"errors"
	"math"
)

// SquareDouble returns the amount of grains on a given square of the chess board
func SquareDouble(input int) (uint64, error) {
	if input <= 0 || input > 64 {
		err := errors.New("Invalid input")
		return 0, err
	}

	return uint64(math.Pow(2, float64(input-1))), nil
}

func Total() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		result, _ := SquareDouble(i)
		sum += result
	}
	return sum
}