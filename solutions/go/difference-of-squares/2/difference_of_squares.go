package diffsquares

func SquareOfSum(n int) int {
	sum_of_integers := n * (n + 1) / 2
	return sum_of_integers * sum_of_integers
}

func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
