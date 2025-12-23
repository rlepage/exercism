package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("invalid square number")
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	return math.MaxUint64 // 2 ** 64 - 1
}
