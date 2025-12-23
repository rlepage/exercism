package prime

import (
	"errors"
	"math"
)

func isPrime(n int) bool {
	if n < 2 || (n%2 == 0 && n != 2) {
		return false
	}
	sqrt := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 3; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return -1, errors.New("Invalid input - should not be less than 1")
	}
	nth := 1
	prime := 2
	test := 3
	for nth < n {
		if isPrime((test)) {
			prime = test
			nth++
		}
		test++
	}
	return prime, nil
}
