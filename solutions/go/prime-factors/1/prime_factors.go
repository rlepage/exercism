package prime

func Factors(n int64) []int64 {
	// log.Printf("Factors(%d)\n", n)

	res := []int64{}
	var divisor int64 = 2
	for n%2 == 0 {
		res = append(res, 2)
		n /= 2
	}
	divisor++

	for n > 1 {
		for n%divisor == 0 {
			res = append(res, int64(divisor))
			n /= int64(divisor)
			// log.Printf("> %v\n", res)
		}
		divisor += 2
	}

	return res
}
