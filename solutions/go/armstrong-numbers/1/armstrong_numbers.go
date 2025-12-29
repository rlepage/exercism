package armstrong

import (
	"strconv"
)

func power(base int, exp int) int {
	var res int = 1
	for exp > 0 {
		res *= base
		exp--
	}
	return res
}

func IsNumber(n int) bool {
	var pow int = len(strconv.Itoa(n))
	var x int = n // to not update n value
	var sum int = 0

	for x > 0 {
		sum += power(x%10, pow)
		x /= 10
	}

	return n == sum
}
