package distribution

import (
	"testing"
)

func TestZipf(t *testing.T) {
	z := NewZipf(1000, 0.693)
	if z.Pdf(1) != z.Cdf(1) {
		t.Error("A values of Pdf(1) and Cdf(1) do not match.")
	}
	if round(z.Pdf(1), 15) != 0.040888100707248 {
		t.Error("A values of Pdf(1) do not match.")
	}
	if round(z.Pdf(1000), 15) != 0.000340876402345 {
		t.Error("A values of Pdf(1000) do not match.")
	}
	if round(z.Cdf(1000), 12) != 1 {
		t.Error("A values of Cdf(1000) do not match.")
	}
}
