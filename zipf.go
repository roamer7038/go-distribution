package distribution

import (
	"math"
)

type Zipf struct {
	a float64
	n int
	f []float64
}

func NewZipf(n int, alpha float64) *Zipf {
	if n <= 0 {
		panic("error")
	}

	var (
		frequency   []float64
		numerator   float64
		denominator float64
	)

	for i := 1; i <= n; i++ {
		denominator += 1.0 / math.Pow(float64(i), alpha)
	}

	for i := 1; i <= n; i++ {
		numerator = 1.0 / math.Pow(float64(i), alpha)
		frequency = append(frequency, numerator/denominator)
	}

	return &Zipf{
		a: alpha,
		n: n,
		f: frequency,
	}
}

func (self *Zipf) Pdf(rank int) float64 {
	if rank < 0 || self.n < rank {
		panic("error")
	}
	return self.f[rank-1]
}

func (self *Zipf) Cdf(rank int) float64 {
	if rank < 0 || self.n < rank {
		panic("error")
	}
	var cdf float64
	for i := 0; i < rank; i++ {
		cdf += self.f[i]
	}
	return cdf
}
