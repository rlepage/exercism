package darts

func squaredRadius(x float64, y float64) float64 {
	return x*x + y*y
}

func Score(x, y float64) int {
	r2 := squaredRadius(x, y)
	switch {
	case r2 <= 1:
		return 10
	case r2 <= 5*5:
		return 5
	case r2 <= 10*10:
		return 1
	default:
		return 0
	}
}
