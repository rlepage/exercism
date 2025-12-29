package sieve

func Sieve(limit int) []int {
	numbers := make(map[int]bool, limit)
	primes := []int{}

	for i := 2; i <= limit; i++ {
		numbers[i] = true
	}

	for x := 2; x <= limit; x++ {
		if numbers[x] {
			primes = append(primes, x)
			for y := 2 * x; y <= limit; y += x {
				numbers[y] = false
			}
		}
	}

	return primes
}
