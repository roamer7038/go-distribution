package distribution

import (
	"math"
)

type Gamma struct {
	k     float64
	thita float64
	f     []float64
	n     int
}

func NewGamma(n int, k float64, thita float64) *Gamma {
	if n < 0 {
		panic("error")
	}

	var (
		frequency []float64
	)

	cdfMax := cdf(n, k, thita)
	for i := 0; i < n; i++ {
		frequency = append(frequency, pdf(i+1, k, thita)/cdfMax)
	}

	return &Gamma{
		k:     k,
		thita: thita,
		f:     frequency,
		n:     n,
	}
}

func pdf(x int, k float64, thita float64) float64 {
	return math.Pow(float64(x), k-1) / (math.Gamma(k) * math.Pow(thita, k)) * math.Exp(float64(-x)/thita)
}

func cdf(x int, k float64, thita float64) float64 {
	var value float64
	for i := 1; i <= x; i++ {
		value += pdf(i, k, thita)
	}
	return value
}

func (self *Gamma) Pdf(x int) float64 {
	if x <= 0 || self.n < x {
		panic("error")
	}
	return self.f[x-1]
}

func (self *Gamma) Cdf(x int) float64 {
	if x <= 0 || self.n < x {
		panic("error")
	}
	var cdf float64
	for i := 0; i < x; i++ {
		cdf += self.f[i]
	}
	return cdf
}
