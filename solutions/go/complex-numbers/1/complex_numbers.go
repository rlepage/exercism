package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	real      float64
	imaginary float64
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.imaginary
}

// (a + ib) + (c + id) = (a + c) + i(b + d)
func (n1 Number) Add(n2 Number) Number {
	return Number{real: n1.real + n2.real, imaginary: n1.imaginary + n2.imaginary}
}

// (a + ib) - (c + id) = (a - c) + i(b - d)
func (n1 Number) Subtract(n2 Number) Number {
	return Number{real: n1.real - n2.real, imaginary: n1.imaginary - n2.imaginary}
}

// (a + ib) (c + id) = (ac - bd) + i(bc + ad)
func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		real:      n1.real*n2.real - n1.imaginary*n2.imaginary,
		imaginary: n1.imaginary*n2.real + n1.real*n2.imaginary,
	}
}

// k (a + ib) = ka + ikb
func (n Number) Times(factor float64) Number {
	return Number{
		real:      factor * n.real,
		imaginary: factor * n.imaginary,
	}
}

// (a + ib) / (c + id) = (ac + bd) / (c^2 + d^2) + i (bc - ad) / (c^2 + d^2)
func (n1 Number) Divide(n2 Number) Number {
	return Number{
		real:      (n1.real*n2.real + n1.imaginary*n2.imaginary) / (n2.real*n2.real + n2.imaginary*n2.imaginary),
		imaginary: (n1.imaginary*n2.real - n1.real*n2.imaginary) / (n2.real*n2.real + n2.imaginary*n2.imaginary),
	}
}

// a - bi
func (n Number) Conjugate() Number {
	return Number{
		real:      n.real,
		imaginary: -n.imaginary,
	}

}

// |z| = sqrt(a^2 + b^2)
func (n Number) Abs() float64 {
	return math.Sqrt(n.real*n.real + n.imaginary*n.imaginary)
}

// e^(a + ib) = e^a e^(ib)
// e^(ib) = cos(b) + i sin(b)
func (n Number) Exp() Number {
	return Number{
		real:      math.Cos(n.imaginary),
		imaginary: math.Sin(n.imaginary),
	}.Times(math.Pow(math.E, n.real))
}
