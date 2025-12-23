package diffsquares

func SquareOfSum(n int) int {
	y := 0
	for i := 1; i <= n; i++ {
		y += i
	}
	return y * y
}

func SumOfSquares(n int) int {
	y := 0
	for i := 1; i <= n; i++ {
		y += i * i
	}
	return y
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
