package distribution

import (
	"testing"
)

func TestGamma(t *testing.T) {
	g := NewGamma(1000, 0.475, 170.607)
	if g.Pdf(1) != g.Cdf(1) {
		t.Error("A values of Pdf(1) and Cdf(1) do not match.")
	}
	if round(g.Pdf(1), 6) != 0.050099 {
		t.Error("A values of Pdf(1) do not match.")
	}
	if round((g.Pdf(1000)*100), 6) != 0.000382 {
		t.Error("A values of Pdf(1000) do not match.")
	}
	if round(g.Cdf(1000), 6) != 1 {
		t.Error("A values of Cdf(1000) do not match.")
	}
}
